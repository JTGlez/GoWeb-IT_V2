package server

func WithPort(port string) func(*Server) error {
	return func(s *Server) error {
		s.port = port
		return nil
	}
}

func WithHost(host string) func(*Server) error {
	return func(s *Server) error {
		s.host = host
		return nil
	}
}
