package telemetry

import (
	"context"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
)

func newTraceExporter() (trace.SpanExporter, error) {
	// ---------- EXPORT TO JAEGER -----------
	headers := map[string]string{
		"content-type": "application/json",
	}
	return otlptrace.New(
		context.Background(),
		otlptracehttp.NewClient(
			otlptracehttp.WithEndpoint("localhost:4318"),
			otlptracehttp.WithHeaders(headers),
			otlptracehttp.WithInsecure(),
		),
	)
	// ---------- EXPORT TO STDOUT -----------
	// return stdoutmetric.New()
}

func newMetricExporter() (metric.Exporter, error) {
	return stdoutmetric.New()
}
