package api

func (s *server) middlewares() {
	s.router.Use(s.metricsMiddleware())
	s.router.Use(s.tracingMiddleware)
	s.router.Use(s.loggingMiddleware)
	// uncommment to use authentication with JSON Wek Keys
	//s.router.Use(s.authenticationMiddleware("test"))
}
