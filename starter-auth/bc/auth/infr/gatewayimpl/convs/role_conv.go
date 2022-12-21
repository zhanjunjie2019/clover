package convs

import (
	"github.com/samber/lo"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/domain/model"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/infr/repo/po"
)

func BatchRolePOToDO(pos []po.Role) []model.Role {
	return lo.Map(pos, func(item po.Role, index int) model.Role {
		return RolePOToDO(item)
	})
}

func RolePOToDO(po po.Role) model.Role {
	return model.NewRole(po.ID, model.RoleValue{
		RoleName: po.RoleName,
		RoleCode: po.RoleCode,
	})
}

func RoleDOToPO(do model.Role) po.Role {
	value := do.FullValue()
	return po.Role{
		ModelPO: defs.ModelPO{
			ID: do.ID(),
		},
		RoleName: value.RoleName,
		RoleCode: value.RoleCode,
	}
}

func RolePermissionPOToValue(po po.RolePermissionRel) model.PermissionValue {
	return model.PermissionValue{
		PermissionName: po.PermissionName,
		AuthCode:       po.AuthCode,
	}
}

func RolePermissionDOToPOs(do model.Role) (rels []po.RolePermissionRel) {
	value := do.FullValue()
	for _, val := range do.GetPermissionValues() {
		rels = append(rels, po.RolePermissionRel{
			RoleId:         do.ID(),
			RoleCode:       value.RoleCode,
			PermissionName: val.PermissionName,
			AuthCode:       val.AuthCode,
		})
	}
	return
}
