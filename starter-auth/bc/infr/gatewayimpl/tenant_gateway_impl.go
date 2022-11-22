package gatewayimpl

import (
	"context"
	"github.com/gogo/protobuf/proto"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/nsqd"
	"github.com/zhanjunjie2019/clover/share/auth/protobuf"
	"github.com/zhanjunjie2019/clover/share/auth/topic"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/domain/model"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/infr/gatewayimpl/convs"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/infr/repo"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type TenantGateway struct {
	TenantRepo  repo.TenantRepoIOCInterface  `singleton:""`
	UserRepo    repo.UserRepoIOCInterface    `singleton:""`
	RoleRepo    repo.RoleRepoIOCInterface    `singleton:""`
	NsqProducer nsqd.NsqProducerIOCInterface `singleton:""`
}

func (t *TenantGateway) FindByTenantID(ctx context.Context, tenantID string) (tenant model.Tenant, exist bool, err error) {
	tenantPO, exist, err := t.TenantRepo.FindByTenantID(ctx, tenantID)
	if err != nil {
		return nil, false, err
	}
	if exist {
		tenant = convs.TenantPOToDO(tenantPO)
	}
	return
}

func (t *TenantGateway) Save(ctx context.Context, tenant model.Tenant) (defs.ID, error) {
	return t.TenantRepo.Save(ctx, convs.TenantDOToPO(tenant))
}

func (t *TenantGateway) PublishInitEvent(ctx context.Context, dto protobuf.TenantInitEventDTO) error {
	bs, err := proto.Marshal(&dto)
	if err != nil {
		return err
	}
	return t.NsqProducer.Publish(ctx, topic.TenantInitTopic, bs)
}

func (t *TenantGateway) TenantTablesManualMigrate(ctx context.Context) (err error) {
	err = t.UserRepo.ManualMigrate(ctx)
	if err != nil {
		return
	}
	err = t.RoleRepo.ManualMigrate(ctx)
	return
}
