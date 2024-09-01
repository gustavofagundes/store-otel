package telemetry

import (
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
)

func newTraceExporter() (trace.SpanExporter, error) {
	return stdouttrace.New()
}

func newMetricExporter() (metric.Exporter, error) {
	return stdoutmetric.New()
}
