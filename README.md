# go-service-template

An example Go service.

## Features

### HTTP API

An HTTP API that uses `gorilla/mux`. It comes with the following built-ins:

  - a `/_system/health` health endpoint;
  - automatic deserialization of opentracing headers on incoming requests, if present. The trace ID, Span ID and parent ID are automatically added to all following logs;
  - automatic log decoration with the route name and HTTP method;
  - (optional) automatic JWT token verification. Note that this does not block requests, only passed authentication information down to the next handler.
  
### Logs

Automatically registers a `uber/zap` logger and exposes it via `logging.GetLogger(ctx). Your logs will contain the fields described above.

### Tracing

Automatically creates an opentracing `Tracer`.

