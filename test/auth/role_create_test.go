package auth

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/zhanjunjie2019/clover/test/auth/consts"
	"github.com/zhanjunjie2019/clover/test/utils"
	"net/http"
	"testing"
	"time"
)

func TestRoleCreate(t *testing.T) {
	Convey("执行[RoleCreate]接口测试过程", t, func() {
		// 超管Token获取
		sadminAccessToken := GetSadminToken()
		So(sadminAccessToken, ShouldNotBeEmpty)

		// 创建租户
		tenantID, secretKey := TenantCreate(sadminAccessToken)
		So(tenantID, ShouldNotBeEmpty)
		So(secretKey, ShouldNotBeEmpty)

		// 租户管理员Token获取
		accessToken := GetTestTenantToken(tenantID, secretKey)
		So(accessToken, ShouldNotBeEmpty)

		time.Sleep(time.Second)

		// 创建角色
		RoleCreate(tenantID, accessToken)
	})
}

func RoleCreate(tenantID, accessToken string) string {
	roleCode := utils.TinyUUID()
	reqBody := map[string]any{
		"roles": []map[string]any{
			{
				"roleName": utils.TinyUUID(),
				"roleCode": roleCode,
			},
			{
				"roleName": utils.TinyUUID(),
				"roleCode": utils.TinyUUID(),
			},
		},
	}
	rs := map[string]any{}
	err := utils.PostRequest(
		consts.DomainHost+consts.RoleCreateApiUrl,
		http.Header{
			"Tenant-ID": []string{tenantID},
			"C-Token":   []string{accessToken},
		},
		reqBody,
		&rs,
	)

	So(err, ShouldBeNil)
	So(rs, ShouldNotBeEmpty)
	So(rs["code"], ShouldEqual, 0)

	data := rs["data"].(map[string]any)
	roleIDs := data["roleIDs"].([]any)
	So(len(roleIDs), ShouldEqual, 2)
	return roleCode
}

func GetTestTenantToken(tenantID, secretKey string) string {
	// 租户登录
	reqBody := map[string]any{
		"tenantID":  tenantID,
		"secretKey": secretKey,
	}
	rs := map[string]any{}
	err := utils.PostRequest(
		consts.DomainHost+consts.TenantTokenCreateApiUrl,
		nil,
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
