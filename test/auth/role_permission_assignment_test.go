package auth

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/zhanjunjie2019/clover/test/auth/consts"
	"github.com/zhanjunjie2019/clover/test/utils"
	"net/http"
	"testing"
)

func TestRolePermissionAssignment(t *testing.T) {

}

func RolePermissionAssignment(tenantID, accessToken, roleCode, authCode string) {
	reqBody := map[string]any{
		"roleCode":  roleCode,
		"authCodes": []string{authCode},
	}
	rs := map[string]any{}
	err := utils.PostRequest(
		consts.DomainHost+consts.RolePermissionAssignmentApiUrl,
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
