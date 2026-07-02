package main

import (
	"log"
	"net/http"
	"raft-kv/internal/kvstore"
	"raft-kv/internal/server"
)

func main() {
	store := kvstore.New()
	srv := server.New(store)

	mux := http.NewServeMux()
	srv.RegisterRoutes(mux)

	log.Println("listening on :9000")
	if err := http.ListenAndServe(":9000", mux); err != nil {
		log.Fatal(err)
	}
}
