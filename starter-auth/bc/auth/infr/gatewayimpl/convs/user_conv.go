package convs

import (
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/domain/model"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/infr/repo/po"
)

func BathUserPOToDO(pos []po.User) (dos []model.User) {
	for i := range pos {
		dos = append(dos, UserPOToDO(pos[i]))
	}
	return
}

func UserPOToDO(po po.User) model.User {
	return model.NewUser(po.ID, model.UserValue{
		UserName: po.UserName,
		Password: po.Password,
	})
}

func UserDOToPO(do model.User) po.User {
	value := do.FullValue()
	return po.User{
		ModelPO: defs.ModelPO{
			ID: do.ID(),
		},
		UserName: value.UserName,
		Password: value.Password,
	}
}

func UserRoleDOToPO(do model.User) (rels []po.UserRoleRel) {
	value := do.FullValue()
	for _, val := range do.GetRoleValues() {
		rels = append(rels, po.UserRoleRel{
			UserID:   do.ID(),
			UserName: value.UserName,
			RoleName: val.RoleName,
			RoleCode: val.RoleCode,
		})
	}
	return
}
