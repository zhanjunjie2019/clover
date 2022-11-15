package gateway

import (
	"context"
	"github.com/zhanjunjie2019/clover/share/auth/dto"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/domain/model"
)

type ITenantGateway interface {
	FindByTenantID(ctx context.Context, tenantID string) (tenant model.Tenant, exist bool, err error)
	Save(ctx context.Context, tenant model.Tenant) error
	PublishInitEvent(ctx context.Context, dto dto.TenantInitEventDTO) error
	TenantTablesManualMigrate(ctx context.Context) error
}
