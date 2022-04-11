package trace

import (
	"context"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	oteltrace "go.opentelemetry.io/otel/trace"

	"github.com/goharbor/harbor/src/lib/log"
	"github.com/goharbor/harbor/src/pkg/version"
)

func initExporter(ctx context.Context) (tracesdk.SpanExporter, error) {
	var err error
	var exp tracesdk.SpanExporter
	cfg := GetGlobalConfig()
	if len(cfg) != 0 {
		
	}
}