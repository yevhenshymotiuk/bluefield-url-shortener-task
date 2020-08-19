package main

func (s *server) routes() {
	s.router.GET("/:id", s.cache(s.handleShortenedURL()))
	s.router.GET("/", s.handleIndex())
	s.router.POST("/", s.handleIndex())
}
