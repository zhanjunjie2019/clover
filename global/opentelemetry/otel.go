package opentelemetry

import (
	"context"
	"fmt"
	"github.com/MrAlias/flow"
	"github.com/gin-gonic/gin"
	"github.com/zhanjunjie2019/clover/global/confs"
	"github.com/zhanjunjie2019/clover/global/consts"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.9.0"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"net/http"
	"time"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type OpenTelemetry struct {
	tracerProvider *sdktrace.TracerProvider
	tracer         trace.Tracer
}

func (o *OpenTelemetry) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := o.tracerProvider.Shutdown(ctx); err != nil {
		otel.Handle(err)
	}
}

func (o *OpenTelemetry) InitProvider() error {
	otelConfig := confs.GetGlobalConfig().OtelConfig
	if otelConfig.Enabled == 1 {
		svcConf := confs.GetServerConfig().SvcConf
		var (
			grpcEndpoint = otelConfig.CollectorGrpcEndpoint
			svcName      = svcConf.SvcName
			svcVersion   = svcConf.SvcVersion
			svcNum       = int(svcConf.SvcNum)
			ctx          = context.Background()
		)

		traceClient := otlptracegrpc.NewClient(
			otlptracegrpc.WithInsecure(),
			otlptracegrpc.WithEndpoint(grpcEndpoint),
			otlptracegrpc.WithDialOption(grpc.WithBlock()))

		ctx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()

		traceExporter, err := otlptrace.New(ctx, traceClient)
		if err != nil {
			return fmt.Errorf("failed to create trace exporter: %w", err)
		}

		res, err := resource.New(ctx,
			resource.WithAttributes(
				semconv.ServiceNameKey.String(svcName),
				semconv.ServiceInstanceIDKey.Int(svcNum),
				semconv.ServiceVersionKey.String(svcVersion),
			),
		)
		if err != nil {
			return fmt.Errorf("failed to create resource: %w", err)
		}

		bsp := sdktrace.NewBatchSpanProcessor(traceExporter)
		tracerProvider := sdktrace.NewTracerProvider(
			sdktrace.WithSampler(sdktrace.AlwaysSample()),
			sdktrace.WithResource(res),
			//sdktrace.WithSpanProcessor(bsp),
			flow.WithSpanProcessor(bsp),
		)

		otel.SetTracerProvider(tracerProvider)
		otel.SetTextMapPropagator(propagation.TraceContext{})
		if o.tracerProvider != nil {
			if err = o.tracerProvider.Shutdown(ctx); err != nil {
				otel.Handle(err)
			}
		}
		o.tracerProvider = tracerProvider
		o.tracer = otel.Tracer("")
	}
	return nil
}

func (o *OpenTelemetry) Start(ctx context.Context, spanName string) (context.Context, trace.Span) {
	return o.tracer.Start(ctx, spanName)
}

func (o *OpenTelemetry) InvokeFunc(ctx context.Context, spanName string, fn func(context.Context, any) (any, error), args any) (any, error) {
	_, ok := ctx.(*gin.Context)
	if ok {
		ctx = uctx.GetSpanContext(ctx)
	}
	if o.tracer != nil {
		c, span := o.tracer.Start(ctx, spanName)
		defer span.End()
		return fn(c, args)
	} else {
		return fn(ctx, args)
	}
}

func (o *OpenTelemetry) Inject(ctx context.Context, headers http.Header) {
	sc := trace.SpanContextFromContext(ctx)
	if !sc.IsValid() {
		return
	}
	headers.Set(consts.TraceIDHeaderKey, sc.TraceID().String())
	headers.Set(consts.TraceSpanIDHeaderKey, sc.SpanID().String())
}

func (o *OpenTelemetry) Extract(ctx context.Context, headers http.Header) context.Context {
	traceID := headers.Get(consts.TraceIDHeaderKey)
	if len(traceID) != 32 {
		return ctx
	}
	spanID := headers.Get(consts.TraceSpanIDHeaderKey)
	if len(spanID) != 16 {
		return ctx
	}
	return o.GetCtx(ctx, traceID, spanID)
}

func (o *OpenTelemetry) GetCtx(ctx context.Context, traceID, spanID string) context.Context {
	var (
		err error
		scc trace.SpanContextConfig
	)
	scc.TraceID, err = trace.TraceIDFromHex(traceID)
	if err != nil {
		return ctx
	}
	scc.SpanID, err = trace.SpanIDFromHex(spanID)
	if err != nil {
		return ctx
	}
	scc.Remote = true

	sc := trace.NewSpanContext(scc)
	if !sc.IsValid() {
		return ctx
	}
	return trace.ContextWithRemoteSpanContext(ctx, sc)
}
