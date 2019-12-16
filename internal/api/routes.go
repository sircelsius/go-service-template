package api

// routes registers all the routes that your HTTP server should handle
func (s *server) routes() {
	// in a production environment, you probably do not want to expose _system endpoints to the world
	s.router.Handle("/_system/health", s.healthHandler()).Name("system/health")
	s.router.Handle("/_system/metrics", s.metricsHandler()).Name("system/metrics")

	s.router.Handle("/cool", s.foaasHandler()).Name("cool")
}
