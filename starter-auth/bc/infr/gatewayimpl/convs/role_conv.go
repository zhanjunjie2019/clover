package convs

import (
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/domain/model"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/infr/repo/po"
)

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
