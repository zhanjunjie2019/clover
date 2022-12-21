package auth

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/zhanjunjie2019/clover/test/auth/consts"
	"github.com/zhanjunjie2019/clover/test/utils"
	"net/http"
	"testing"
	"time"
)

func TestUserTokenByAuthCode(t *testing.T) {
	Convey("执行[UserTokenByAuthCode]接口测试过程", t, func() {
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
		tenantAccessToken := GetTestTenantToken(tenantID, secretKey)
		So(tenantAccessToken, ShouldNotBeEmpty)

		time.Sleep(time.Second)

		// 创建角色
		roleCode := RoleCreate(tenantID, tenantAccessToken)
		So(roleCode, ShouldNotBeEmpty)

		// 角色关联资源
		RolePermissionAssignment(tenantID, tenantAccessToken, roleCode, authCode)

		// 创建用户
		username, password := UserCreate(tenantID, tenantAccessToken)
		So(username, ShouldNotBeEmpty)
		So(password, ShouldNotBeEmpty)

		// 用户关联角色
		UserRoleAssignment(tenantID, tenantAccessToken, username, roleCode)

		// 获取用户授权码
		authorizationCode := UserAuthorizationCode(tenantID, username, password)
		So(authorizationCode, ShouldNotBeEmpty)

		// 通过授权码获取用户token
		userAccessToken := UserTokenByAuthCode(tenantID, authorizationCode)
		So(userAccessToken, ShouldNotBeEmpty)
		fmt.Println(userAccessToken)
	})
}

func UserTokenByAuthCode(tenantID, authCode string) string {
	reqBody := map[string]any{
		"authorizationCode": authCode,
	}
	rs := map[string]any{}
	err := utils.PostRequest(
		consts.DomainHost+consts.UserTokenByAuthCodeApiUrl,
		http.Header{
			"Tenant-ID": []string{tenantID},
		},
		reqBody,
		&rs,
	)

	So(err, ShouldBeNil)
	So(rs, ShouldNotBeEmpty)
	So(rs["code"], ShouldEqual, 0)

	data := rs["data"].(map[string]any)
	accessToken := data["accessToken"].(string)
	So(accessToken, ShouldNotBeEmpty)
	return accessToken
}
