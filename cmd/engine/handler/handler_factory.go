package handler

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/taskifyworks/api/cmd/app/usecases"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"io"
	"log"
	"net/http"
	"time"
)

type (
	Factory interface {
		CreateRestHandler(gc *GitHubConfig) http.Handler
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

func (hf *handlerFactory) CreateRestHandler(gc *GitHubConfig) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/oauth/csrf", func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, time.Now().String())
	})

	oauth := r.PathPrefix("/oauth/signup").Methods(http.MethodGet).Subrouter()
	oauth.HandleFunc("/github", createGitHubSignUp(gc))
	oauth.HandleFunc("/github/callback", createGitHubCallback(gc))

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000", "https://taskify.works", "https://*.taskify.works"},
	})
	return c.Handler(r)
}

func NewFactory(uf usecases.InteractorFactory) Factory {
	return &handlerFactory{
		UseCaseFactory: uf,
	}
}

func createGitHubSignUp(gc *GitHubConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := &oauth2.Config{
			ClientID:     gc.ClientId,
			ClientSecret: gc.ClientSecret,
			RedirectURL:  "http://localhost:3001/oauth/signup/github/callback",
			Endpoint:     github.Endpoint,
		}

		state := "My_Secret_Code"

		url := c.AuthCodeURL(state)
		fmt.Println(url)
		http.Redirect(w, r, url, http.StatusSeeOther)
	}
}

func createGitHubCallback(gc *GitHubConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := &oauth2.Config{
			ClientID:     gc.ClientId,
			ClientSecret: gc.ClientSecret,
			RedirectURL:  "http://localhost:3001/oauth/signup/github/callback",
			Endpoint:     github.Endpoint,
		}
		//state := r.URL.Query().Get("state")
		code := r.URL.Query().Get("code")
		hc := &http.Client{Timeout: 2 * time.Second}
		ctx := context.WithValue(context.Background(), oauth2.HTTPClient, hc)
		tok, err := c.Exchange(ctx, code)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(tok)
		client := c.Client(ctx, tok)
		_ = client
		http.Redirect(w, r, "https://app.taskify.works", http.StatusSeeOther)
	}
}
