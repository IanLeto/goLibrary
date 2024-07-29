package main

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"goLibrary/utils"
	//trace2 "go.opentelemetry.io/otel/sdk/trace"
)

var traceIan = otel.Tracer("goLib")

func NewTraceProvider(url string) (*trace.TracerProvider, error) {
	traceExporter, err := stdouttrace.New(
		stdouttrace.WithPrettyPrint())
	utils.NoErr(err)
	tp := trace.NewTracerProvider(trace.WithBatcher(traceExporter))
	trace.WithResource(resource.NewWithAttributes(semconv.SchemaURL,
		semconv.ServiceNameKey.String("ian"),
		attribute.String("test", "xxx"),
	))
	return tp, nil
}
