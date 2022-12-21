package auth

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/zhanjunjie2019/clover/test/auth/consts"
	"github.com/zhanjunjie2019/clover/test/utils"
	"net/http"
	"testing"
)

func TestPermissionCreate(t *testing.T) {
	Convey("执行[PermissionCreate]接口测试过程", t, func() {
		// 超管Token获取
		accessToken := GetSadminToken()
		So(accessToken, ShouldNotBeEmpty)

		// 创建资源
		PermissionCreate(accessToken)
	})
}

func PermissionCreate(accessToken string) string {
	authCode := utils.TinyUUID()
	reqBody := map[string]any{
		"permissions": []map[string]any{
			{
				"permissionName": utils.TinyUUID(),
				"authCode":       authCode,
			},
			{
				"permissionName": utils.TinyUUID(),
				"authCode":       utils.TinyUUID(),
			},
		},
	}
	rs := map[string]any{}
	err := utils.PostRequest(
		consts.DomainHost+consts.PermissionCreateApiUrl,
		http.Header{
			"C-Token": []string{
				accessToken,
			},
		},
		reqBody,
		&rs,
	)

	So(err, ShouldBeNil)
	So(rs, ShouldNotBeEmpty)
	So(rs["code"], ShouldEqual, 0)

	data := rs["data"].(map[string]any)
	permissionIDs := data["permissionIDs"].([]any)
	So(len(permissionIDs), ShouldEqual, 2)
	return authCode
}

func GetSadminToken() string {
	reqBody := map[string]any{
		"secretKey": "CloverSecretKeyCloverSecretKey",
	}
	rs := map[string]any{}

	err := utils.PostRequest(
		consts.DomainHost+consts.SadminTokenCreateApiUrl,
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
