package tracing

import (
	"context"
	"io"

	"github.com/sircelsius/go-service-template/internal/logging"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-client-go/log/zap"
	"github.com/uber/jaeger-lib/metrics"
)

func NewTracer(ctx context.Context, serviceName string, development bool) io.Closer {
	var cfg config.Configuration
	if development {
		cfg = config.Configuration{
			ServiceName:         serviceName,
			Sampler: &config.SamplerConfig{
				Type:  jaeger.SamplerTypeConst,
				Param: 1,
			},
			Reporter: &config.ReporterConfig{
				LogSpans: true,
			},
		}
	} else {
		cfg = config.Configuration{}
	}
	logger := logging.GetLogger(ctx)

	closer, err := cfg.InitGlobalTracer(
		serviceName,
		config.Logger(zap.NewLogger(logger)),
		config.Metrics(metrics.NullFactory),
	)

	if err != nil {
		panic(err)
	}
	return closer
}