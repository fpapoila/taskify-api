package handler

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/taskifyworks/api/cmd/app/usecases"
	"net/http"
)

type (
	Factory interface {
		CreateRestHandler() http.Handler
	}

	handlerFactory struct {
		UseCaseFactory usecases.InteractorFactory
	}

	GitHubConfig struct {
		AppID         string
		PrivateKey    string
		WebHookSecret string
		ClientId      string
		ClientSecret  string
	}
)

func (hf *handlerFactory) CreateRestHandler() http.Handler {
	r := mux.NewRouter()

	gh := hf.UseCaseFactory.CreateGitHubInteractor()
	r.HandleFunc("/oauth/signup/github", func(w http.ResponseWriter, r *http.Request) {

		gh.SignUp()
	})

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000", "https://*.taskify.works"},
	})
	return c.Handler(r)
}

func NewFactory(uf usecases.InteractorFactory) Factory {
	return &handlerFactory{
		UseCaseFactory: uf,
	}
}
