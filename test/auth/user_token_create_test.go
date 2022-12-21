package auth

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestUserTokenCreate(t *testing.T) {
	Convey("执行[UserTokenCreate]接口测试过程", t, func() {
		// 超管Token获取
		sadminAccessToken := GetSadminToken()
		So(sadminAccessToken, ShouldNotBeEmpty)

		// 创建租户
		tenantID, secretKey := TenantCreate(sadminAccessToken)
		So(tenantID, ShouldNotBeEmpty)
		So(secretKey, ShouldNotBeEmpty)

		// 创建资源
		authCode := PermissionCreate(sadminAccessToken)
		So(authCode, ShouldNotBeEmpty)

		// 租户管理员Token获取
		accessToken := GetTestTenantToken(tenantID, secretKey)
		So(accessToken, ShouldNotBeEmpty)

		time.Sleep(time.Second)

		// 创建角色
		roleCode := RoleCreate(tenantID, accessToken)
		So(roleCode, ShouldNotBeEmpty)

		// 角色关联资源
		RolePermissionAssignment(tenantID, accessToken, roleCode, authCode)

		// 创建用户
		username, password := UserCreate(tenantID, accessToken)
		So(username, ShouldNotBeEmpty)
		So(password, ShouldNotBeEmpty)

		// 用户关联角色
		UserRoleAssignment(tenantID, accessToken, username, roleCode)

		// 获取用户授权码
		authorizationCode := UserAuthorizationCode(tenantID, username, password)
		So(authorizationCode, ShouldNotBeEmpty)

		// 通过授权码获取用户token
	})
}
