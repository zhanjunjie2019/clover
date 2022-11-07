//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to re-generate

package opentelemetry

import (
	contextx "context"
	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	singleton "github.com/alibaba/ioc-golang/autowire/singleton"
	util "github.com/alibaba/ioc-golang/autowire/util"
	"go.opentelemetry.io/otel/trace"
	"net/http"
)

func init() {
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &openTelemetry_{}
		},
	})
	openTelemetryStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &OpenTelemetry{}
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
	}
	singleton.RegisterStructDescriptor(openTelemetryStructDescriptor)
}

type openTelemetry_ struct {
	Shutdown_     func()
	InitProvider_ func() error
	Start_        func(ctx contextx.Context, spanName string) (contextx.Context, trace.Span)
	InvokeFunc_   func(ctx contextx.Context, spanName string, fn func(contextx.Context, any) (any, error), args any) (any, error)
	Inject_       func(ctx contextx.Context, headers http.Header)
	Extract_      func(ctx contextx.Context, headers http.Header) contextx.Context
	GetCtx_       func(ctx contextx.Context, traceID, spanID string) contextx.Context
}

func (o *openTelemetry_) Shutdown() {
	o.Shutdown_()
}

func (o *openTelemetry_) InitProvider() error {
	return o.InitProvider_()
}

func (o *openTelemetry_) Start(ctx contextx.Context, spanName string) (contextx.Context, trace.Span) {
	return o.Start_(ctx, spanName)
}

func (o *openTelemetry_) InvokeFunc(ctx contextx.Context, spanName string, fn func(contextx.Context, any) (any, error), args any) (any, error) {
	return o.InvokeFunc_(ctx, spanName, fn, args)
}

func (o *openTelemetry_) Inject(ctx contextx.Context, headers http.Header) {
	o.Inject_(ctx, headers)
}

func (o *openTelemetry_) Extract(ctx contextx.Context, headers http.Header) contextx.Context {
	return o.Extract_(ctx, headers)
}

func (o *openTelemetry_) GetCtx(ctx contextx.Context, traceID, spanID string) contextx.Context {
	return o.GetCtx_(ctx, traceID, spanID)
}

type OpenTelemetryIOCInterface interface {
	Shutdown()
	InitProvider() error
	Start(ctx contextx.Context, spanName string) (contextx.Context, trace.Span)
	InvokeFunc(ctx contextx.Context, spanName string, fn func(contextx.Context, any) (any, error), args any) (any, error)
	Inject(ctx contextx.Context, headers http.Header)
	Extract(ctx contextx.Context, headers http.Header) contextx.Context
	GetCtx(ctx contextx.Context, traceID, spanID string) contextx.Context
}

var _openTelemetrySDID string

func GetOpenTelemetrySingleton() (*OpenTelemetry, error) {
	if _openTelemetrySDID == "" {
		_openTelemetrySDID = util.GetSDIDByStructPtr(new(OpenTelemetry))
	}
	i, err := singleton.GetImpl(_openTelemetrySDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*OpenTelemetry)
	return impl, nil
}

func GetOpenTelemetryIOCInterfaceSingleton() (OpenTelemetryIOCInterface, error) {
	if _openTelemetrySDID == "" {
		_openTelemetrySDID = util.GetSDIDByStructPtr(new(OpenTelemetry))
	}
	i, err := singleton.GetImplWithProxy(_openTelemetrySDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(OpenTelemetryIOCInterface)
	return impl, nil
}

type ThisOpenTelemetry struct {
}

func (t *ThisOpenTelemetry) This() OpenTelemetryIOCInterface {
	thisPtr, _ := GetOpenTelemetryIOCInterfaceSingleton()
	return thisPtr
}
