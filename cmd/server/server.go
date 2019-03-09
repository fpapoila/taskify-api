package server

import (
	"github.com/taskifyworks/api/cmd/server/handler"
	"net/http"
)

type (
	Server interface {
		Run() error
	}
	server struct {
		port           string
		handlerFactory handler.Factory
	}
)

func New() Server {
	c := GetConfig()
	return &server{
		port:           c.Port,
		handlerFactory: handler.NewFactory(c.Github),
	}
}

func (s *server) Run() error {
	return http.ListenAndServe(":"+s.port, s.handlerFactory.CreateRestHandler())
}
