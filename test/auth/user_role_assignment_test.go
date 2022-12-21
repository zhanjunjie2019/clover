package auth

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/zhanjunjie2019/clover/test/auth/consts"
	"github.com/zhanjunjie2019/clover/test/utils"
	"net/http"
	"testing"
)

func TestUserRoleAssignment(t *testing.T) {

}

func UserRoleAssignment(tenantID, accessToken, username, roleCode string) {
	reqBody := map[string]any{
		"userName":  username,
		"roleCodes": []string{roleCode},
	}
	rs := map[string]any{}
	err := utils.PostRequest(
		consts.DomainHost+consts.UserRoleAssignmentApiUrl,
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
}
