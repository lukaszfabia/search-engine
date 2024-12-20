package server

import (
	"net/http"

	"search-engine/cmd/web"

	"github.com/gorilla/mux"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/", web.HomeHandler)

	r.HandleFunc("/match", web.OnChangeSearchHandler)

	r.HandleFunc("/search", web.OnSubmitSearchHandler)

	r.HandleFunc("/add", web.AddElemHandler)

	fileServer := http.FileServer(http.FS(web.Files))
	r.PathPrefix("/assets/").Handler(fileServer)

	return r
}
