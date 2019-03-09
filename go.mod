module github.com/taskifyworks/api

// +heroku goVersion go1.12
// +heroku install ./cmd/...

require (
	github.com/gorilla/mux v1.7.0
	github.com/rs/cors v1.6.0
	golang.org/x/oauth2 v0.0.0-20190226205417-e64efc72b421
)
