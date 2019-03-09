package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/taskifyworks/api/cmd/server"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, "Hello API")
	})
	r.HandleFunc("/oauth/signup/github", func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, "login")
	})

	h := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000", "https://*.taskify.works"},
	}).Handler(r)
	log.Fatal(http.ListenAndServe(":"+port, h))

	s := server.New()
	log.Fatal(s.Run())

}
