package convs

import (
	"github.com/samber/lo"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/domain/model"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/infr/repo/po"
)

func BatchPermissionPOToDO(pos []po.Permission) []model.Permission {
	return lo.Map(pos, func(item po.Permission, index int) model.Permission {
		return PermissionPOToDO(item)
	})
}

func PermissionPOToDO(po po.Permission) model.Permission {
	return model.NewPermission(po.ID, model.PermissionValue{
		PermissionName: po.PermissionName,
		AuthCode:       po.AuthCode,
	})
}

func PermissionDOToPO(do model.Permission) po.Permission {
	value := do.FullValue()
	return po.Permission{
		ModelPO: defs.ModelPO{
			ID: do.ID(),
		},
		PermissionName: value.PermissionName,
		AuthCode:       value.AuthCode,
	}
}
