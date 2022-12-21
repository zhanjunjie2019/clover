package auth

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/zhanjunjie2019/clover/test/auth/consts"
	"github.com/zhanjunjie2019/clover/test/utils"
	"net/http"
	"testing"
)

func TestUserAuthorizationCode(t *testing.T) {

}

func UserAuthorizationCode(tenantID, username, password string) string {
	reqBody := map[string]any{
		"userName": username,
		"password": password,
	}
	rs := map[string]any{}
	err := utils.PostRequest(
		consts.DomainHost+consts.UserAuthorizationCodeApiUrl,
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
	authorizationCode := data["authorizationCode"].(string)
	So(authorizationCode, ShouldNotBeEmpty)
	return authorizationCode
}
