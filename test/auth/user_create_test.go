package auth

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/zhanjunjie2019/clover/test/auth/consts"
	"github.com/zhanjunjie2019/clover/test/utils"
	"net/http"
	"testing"
)

func TestUserCreate(t *testing.T) {

}

func UserCreate(tenantID, accessToken string) (username, password string) {
	username = utils.TinyUUID()
	password = utils.TinyUUID()
	reqBody := map[string]any{
		"users": []map[string]any{
			{
				"userName": username,
				"password": password,
			},
			{
				"userName": utils.TinyUUID(),
				"password": utils.TinyUUID(),
			},
		},
	}
	rs := map[string]any{}
	err := utils.PostRequest(
		consts.DomainHost+consts.UserCreateApiUrl,
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
	return
}
