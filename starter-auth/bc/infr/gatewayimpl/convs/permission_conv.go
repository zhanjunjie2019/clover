package convs

import (
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/domain/model"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/infr/repo/po"
)

func BatchPermissionPOToDO(pos []po.Permission) (dos []model.Permission) {
	for i := range pos {
		dos = append(dos, PermissionPOToDO(pos[i]))
	}
	return
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
