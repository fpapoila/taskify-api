package handler

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"log"
	"net/http"
	"time"
)

type (
	GitHubHandlerFactory interface {
		CreateSignUpHandler() http.HandlerFunc
		CreateSignUpCallbackHandler() http.HandlerFunc
		CreateWebHookHandler() http.HandlerFunc
	}

	gitHubHandlerFactory struct {
		config *GitHubConfig
	}
)

func (gh *gitHubHandlerFactory) CreateSignUpHandler() http.HandlerFunc {
	return nil
}

func NewGitHubHandlerFactory(c *GitHubConfig) GitHubHandlerFactory {
	return nil
}

func GitHubSignUp(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		Scopes:       nil,
		Endpoint:     github.Endpoint,
	}
	_ = conf.AuthCodeURL("", oauth2.AccessTypeOffline)
	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}
	httpClient := &http.Client{Timeout: 2 * time.Second}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient)

	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}

	client := conf.Client(ctx, tok)
	_ = client
}
