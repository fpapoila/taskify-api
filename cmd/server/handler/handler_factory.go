package handler

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
)

type (
	Factory interface {
		CreateRestHandler() http.Handler
	}

	handlerFactory struct {
		gitHubConfig *GitHubConfig
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

	gh := NewGitHubHandlerFactory(hf.gitHubConfig)
	r.HandleFunc("/oauth/signup/github", gh.CreateSignUpHandler()).Methods(http.MethodGet)
	r.HandleFunc("/oauth/signup/github/callback", gh.CreateSignUpCallbackHandler()).Methods(http.MethodPost)
	r.HandleFunc("/webhooks/github", gh.CreateWebHookHandler()).Methods(http.MethodPost)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000", "https://*.taskify.works"},
	})
	return c.Handler(r)
}

func NewFactory(gc *GitHubConfig) Factory {
	return &handlerFactory{
		gitHubConfig: gc,
	}
}
