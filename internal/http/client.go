package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/sony/gobreaker"
)

type client struct {
	client *http.Client
	name string
	cb     *gobreaker.CircuitBreaker
}

func NewClient(name string) *client {
	return &client{
		client: &http.Client{
			Transport: &http.Transport{
				IdleConnTimeout:       time.Minute,
				TLSHandshakeTimeout:   100 * time.Millisecond,
				ResponseHeaderTimeout: 100 * time.Millisecond,
			},
			Timeout: 500 * time.Millisecond,
		},
		cb: gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:          name,
			MaxRequests:   10,
			Interval:      time.Minute,
			Timeout:       5 * time.Second,
			OnStateChange: nil,
		}),
		name: name,
	}
}

func (c *client) Do(ctx context.Context, r *http.Request) (*http.Response, error) {
	r, err := injectTracingHeaders(ctx, r)
	if err != nil {
		return nil, err
	}

	resp, err := c.cb.Execute(func() (interface{}, error) {
		return c.client.Do(r)
	})

	if err != nil {
		if err == gobreaker.ErrOpenState {
			return nil, errors.Wrap(err, fmt.Sprintf("%v circuit breaker is open", c.name))
		}
		if err == gobreaker.ErrTooManyRequests {
			return nil, errors.Wrap(err, fmt.Sprintf("%v circuit breaker is half open", c.name))
		}
		return nil, err
	}

	if response := resp.(*http.Response); resp != nil {
		return response, nil
	}
	return nil, errors.New("unable to cast response to http.Response")
}

func injectTracingHeaders(ctx context.Context, r *http.Request) (*http.Request, error) {
	if span := opentracing.SpanFromContext(ctx); span != nil {
		err := opentracing.GlobalTracer().Inject(
			span.Context(),
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(r.Header))
		if err != nil {
			return nil, errors.Wrap(err, "unable to inject opentracing headers")
		}
	}
	return r, nil
}