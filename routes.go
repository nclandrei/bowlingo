package main

// this is such a short file because now we have only one handler, but in
// the future, it can have multiple ones and it is much easier to follow the
// code path using this method
func (s *Server) routes() {
	s.router.Handle("/api/score", s.ScoreHandler())
}
