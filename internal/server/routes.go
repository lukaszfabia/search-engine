package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gorilla/mux"
	"search-engine/cmd/web"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/", s.HelloWorldHandler)

	fileServer := http.FileServer(http.FS(web.Files))
	r.PathPrefix("/assets/").Handler(fileServer)

	r.HandleFunc("/web", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(web.HelloForm()).ServeHTTP(w, r)
	})

	r.HandleFunc("/hello", web.HelloWebHandler)

	return r
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
