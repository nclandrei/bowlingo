package bowlingo

// this is such a short file because now we have only one handler, but in
// the future, it can have multiple ones and it is much easier to follow the
// code path using this approach.

// Routes attaches all the defined routes for our bowling server.
func (s *Server) Routes() {
	s.router.Handle("/api/score", s.ScoreHandler())
}
