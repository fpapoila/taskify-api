package server

import (
	"errors"
	"fmt"
	"github.com/taskifyworks/api/cmd/server/handler"
	"log"
	"os"
)

type (
	Config struct {
		Port     string
		MongoURI string
		Github   *handler.GitHubConfig
	}

)

func GetConfig() *Config {
	return &Config{
		Port:     must(get("PORT")),
		MongoURI: must(get("MONGODB_URI")),
		Github: &handler.GitHubConfig{
			AppID:         must(get("GITHUB_APP_ID")),
			PrivateKey:    must(get("GITHUB_PRIVATE_KEY")),
			WebHookSecret: must(get("GITHUB_WEBHOOK_SECRET")),
			ClientId:      must(get("GITHUB_CLIENT_ID")),
			ClientSecret:  must(get("GITHUB_CLIENT_SECRET")),
		},
	}
}

func must(value string, err error) string {
	if err != nil {
		log.Panic(err)
	}
	return value
}

func get(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", errors.New(fmt.Sprintf("$%s should not be empty", key))
	}
	return value, nil
}
