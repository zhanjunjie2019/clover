package convs

import (
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/domain/model"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/infr/repo/po"
)

func TenantPOToDO(po po.Tenant) model.Tenant {
	return model.NewTenant(po.ID, model.TenantValue{
		TenantID:              po.TenantID,
		TenantName:            po.TenantName,
		SecretKey:             po.SecretKey,
		RedirectUrl:           po.RedirectUrl,
		AccessTokenTimeLimit:  po.AccessTokenTimeLimit,
		RefreshTokenTimeLimit: po.RefreshTokenTimeLimit,
	})
}

func TenantDOToPO(do model.Tenant) po.Tenant {
	value := do.FullValue()
	return po.Tenant{
		ModelPO: defs.ModelPO{
			ID: do.ID(),
		},
		TenantID:              value.TenantID,
		TenantName:            value.TenantName,
		SecretKey:             value.SecretKey,
		RedirectUrl:           value.RedirectUrl,
		AccessTokenTimeLimit:  value.AccessTokenTimeLimit,
		RefreshTokenTimeLimit: value.RefreshTokenTimeLimit,
	}
}
