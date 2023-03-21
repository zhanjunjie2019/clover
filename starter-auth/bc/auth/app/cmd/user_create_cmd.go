package cmd

import "github.com/zhanjunjie2019/clover/global/defs"

type UserCreateCmd struct {
	Users []UserInfo `json:"users"`
}

type UserInfo struct {
	// 用户名
	UserName string `json:"user_name"`
	// 密码
	Password string `json:"password"`
}

type UserCreateResult struct {
	UserIDs []defs.ID `json:"user_ids"`
}
