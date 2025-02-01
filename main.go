package main

import (
	"log"
	"net/http"
)

type api struct {
	addr string
}

func (a *api) getUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User page bitches"))
}

func (a *api) createUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Created User"))
}

func main() {

	mux := http.NewServeMux()

	api := &api{
		addr: ":8080",
	}

	srv := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.getUserHandler)
	mux.HandleFunc("POST /users", api.createUsersHandler)

	if err := http.ListenAndServe(srv.Addr, srv.Handler); err != nil {
		log.Fatal(err)
	}
}
