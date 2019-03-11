package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/taskifyworks/api/cmd/app/usecases"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

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
	oauth := r.PathPrefix("/oauth").Subrouter()

	oauth.HandleFunc("/github", func(w http.ResponseWriter, r *http.Request) {
		state := GetRandomString(32)
		c := &oauth2.Config{
			ClientID:     gc.ClientId,
			ClientSecret: gc.ClientSecret,
			RedirectURL:  "http://localhost:3000/oauth/signup/github",
			Endpoint:     github.Endpoint,
		}
		url := c.AuthCodeURL(state)

		d := struct {
			CSRFToken string `json:"csrfToken"`
			URL       string `json:"url"`
		}{
			state,
			url,
		}
		_ = json.NewEncoder(w).Encode(d)
	}).Methods(http.MethodGet)
	oauth.HandleFunc("/signup/github", createGitHubSignUp(gc)).Methods(http.MethodPost)

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
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		defer r.Body.Close()
		state := string(b)

		c := &oauth2.Config{
			ClientID:     gc.ClientId,
			ClientSecret: gc.ClientSecret,
			RedirectURL:  "http://localhost:3000/oauth/signup/github",
			Endpoint:     github.Endpoint,
		}

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

func GetRandomString(length uint8) string {
	result := make([]byte, length)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
