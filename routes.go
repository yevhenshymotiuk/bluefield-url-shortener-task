package main

func (s *server) routes() {
	s.router.GET("/:id", s.handleShortenedURL())
	s.router.GET("/", s.handleIndex())
}
