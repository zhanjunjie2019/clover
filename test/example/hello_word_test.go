package example

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/zhanjunjie2019/clover/test/utils"
	"net/http"
	"net/url"
	"testing"
)

const (
	bashUrl          = "http://127.0.0.1:8800"
	helloworldApiUrl = "/example/hellow-world"
)

func TestHelloWordController(t *testing.T) {
	Convey("执行[HelloWorld]接口测试", t, func() {
		qry := url.Values{}
		qry.Add("firstName", utils.UUID())
		qry.Add("lastName", utils.UUID())
		rs := map[string]any{}

		err := utils.GetRequest(
			bashUrl+helloworldApiUrl,
			http.Header{},
			qry,
			&rs,
		)

		So(err, ShouldBeNil)
		So(rs, ShouldNotBeEmpty)
		So(rs["code"], ShouldEqual, 0)
	})
}
