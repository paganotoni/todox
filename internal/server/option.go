package server

type Option func(*Root)

func WithHost(host string) Option {
	return func(s *Root) {
		s.host = host
	}
}

func WithPort(port string) Option {
	return func(s *Root) {
		s.port = port
	}
}
