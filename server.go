package fund_management_information_system

import (
	"net/http"
	"time"
)

const Timeout = 10 * time.Second

type Server struct {
	server *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.server = &http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		ReadTimeout:  Timeout,
		WriteTimeout: Timeout,
	}

	return s.server.ListenAndServe()
}
