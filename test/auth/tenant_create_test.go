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

func TestTenantCreate(t *testing.T) {
	Convey("执行[TenantCreate]接口测试过程", t, func() {
		// 超管Token获取
		accessToken := GetSadminToken()
		So(accessToken, ShouldNotBeEmpty)

		// 创建租户
		TenantCreate(accessToken)
	})
}

func TenantCreate(accessToken string) (tenantID, secretKey string) {
	reqBody := map[string]any{
		"tenants": []map[string]any{
			{
				"tenantID":   fmt.Sprintf("%d", time.Now().Unix()),
				"tenantName": utils.TinyUUID(),
			},
			{
				"tenantName": utils.TinyUUID(),
			},
		},
	}
	rs := map[string]any{}
	err := utils.PostRequest(
		consts.DomainHost+consts.TenantCreateApiUrl,
		http.Header{
			"C-Token": []string{accessToken},
		},
		reqBody,
		&rs,
	)

	So(err, ShouldBeNil)
	So(rs, ShouldNotBeEmpty)
	So(rs["code"], ShouldEqual, 0)

	data := rs["data"].(map[string]any)
	secretKeys := data["secretKeys"].([]any)
	So(len(secretKeys), ShouldEqual, 2)
	tenantID = secretKeys[0].(map[string]any)["tenantID"].(string)
	secretKey = secretKeys[0].(map[string]any)["secretKey"].(string)
	So(tenantID, ShouldNotBeEmpty)
	So(secretKey, ShouldNotBeEmpty)
	return
}
