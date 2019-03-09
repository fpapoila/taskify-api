package main

import (
	"github.com/taskifyworks/api/cmd/engine"
	"log"
)

func main() {
	s := engine.NewServer()
	log.Fatal(s.Run())
}
