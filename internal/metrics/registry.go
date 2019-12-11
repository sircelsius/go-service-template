package metrics

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/graphite"
	"github.com/prometheus/common/log"
)

func StartMetrics(ctx context.Context) {
	b, err := graphite.NewBridge(&graphite.Config{
		URL:           "localhost:2003",
		Gatherer:      prometheus.DefaultGatherer,
		Prefix:        "prefix",
		Interval:      30 * time.Second,
		Timeout:       10 * time.Second,
		ErrorHandling: graphite.AbortOnError,
		Logger:        log.NewErrorLogger(),
	})
	if err != nil {
		panic(err)
	}
	// Push initial metrics to Graphite. Fail fast if the push fails.
	if err := b.Push(); err != nil {
		panic(err)
	}

	// Create a Context to control stopping the Run() loop that pushes
	// metrics to Graphite.
	ctx, _ = context.WithCancel(ctx)

	// Start pushing metrics to Graphite in the Run() loop.
	go b.Run(ctx)
}
