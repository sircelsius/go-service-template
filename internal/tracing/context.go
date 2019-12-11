package tracing

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/sircelsius/go-service-template/internal/logging"
	"github.com/uber/jaeger-client-go"
	"go.uber.org/zap"
)

const (
	traceLoggerKey = "trace_id_s"
	spanLoggerKey = "span_id_s"
	parentLoggerKey = "parent_id_s"
)

// CreateSpan creates a span from the current context and returns the span as well as a context
// with a decorated logger.
func CreateSpan(ctx context.Context, operationName string) (opentracing.Span, context.Context) {
	span, ctx := opentracing.StartSpanFromContext(ctx, operationName)

	ctx = ContextWithSpanLogger(ctx, span)
	return span, ctx
}

// ContextWithSpanLogger adds the information of the given span to the given context's logger.
func ContextWithSpanLogger(ctx context.Context, span opentracing.Span) context.Context {
	var logger *zap.Logger
	if sc, ok := span.Context().(jaeger.SpanContext); ok {
		logger = logging.GetLogger(ctx).With(
			zap.String(traceLoggerKey, sc.TraceID().String()),
			zap.String(spanLoggerKey, sc.SpanID().String()),
			zap.String(parentLoggerKey, sc.ParentID().String()),
		)
		ctx = logging.ContextWithLogger(ctx, logger)
	}
	return ctx
}

