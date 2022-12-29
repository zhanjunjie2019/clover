package grpchandler

import (
	"context"
	"github.com/zhanjunjie2019/clover/share/example/protobuf"
	"github.com/zhanjunjie2019/clover/starter-example/bc/example/app"
	"github.com/zhanjunjie2019/clover/starter-example/bc/example/app/cmd"
	"go-micro.dev/v4/server"
)

// +ioc:autowire=true
// +ioc:autowire:type=allimpls
// +ioc:autowire:implements=github.com/zhanjunjie2019/clover/global/defs.IGrpcHandler

type HelloWorldGrpcHandler struct {
	ExampleApp app.ExampleAppIOCInterface `singleton:""`
}

func (h *HelloWorldGrpcHandler) GrpcRegister(s server.Server) error {
	return protobuf.RegisterExampleServiceHandler(s, h)
}

func (h *HelloWorldGrpcHandler) HelloWorld(ctx context.Context, reqVO *protobuf.ExampleGrpcReqVO, rspVO *protobuf.ExampleGrpcRspVO) error {
	result, err := h.ExampleApp.ExampleHellowWorld(ctx, cmd.HelloWorldCmd{
		FirstName: reqVO.FirstName,
		LastName:  reqVO.LastName,
	})
	if err == nil {
		rspVO.Greetings = result.Greetings
	}
	return err
}
