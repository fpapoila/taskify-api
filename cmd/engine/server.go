package engine

import (
	"github.com/taskifyworks/api/cmd/app/usecases"
	"github.com/taskifyworks/api/cmd/engine/handler"
	"net/http"
)

type (
	Server interface {
		Run() error
	}

	server struct {
		port    string
		handler http.Handler
	}
)

func NewServer() Server {
	c := GetConfig()
	uf := usecases.NewFactory()
	hf := handler.NewFactory(uf)

	return &server{
		port:    c.Port,
		handler: hf.CreateRestHandler(),
	}
}

func (s *server) Run() error {
	return http.ListenAndServe(":"+s.port, s.handler)
}
