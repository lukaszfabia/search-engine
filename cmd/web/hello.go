package web

import (
	"fmt"
	"log"
	"net/http"
	"search-engine/internal/matcher"
)

var engine = matcher.NewMatcher()

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	component := Search()
	err = component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error rendering in HelloWebHandler: %e", err)
	}
}

func AddElemHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	elem := r.FormValue("elem")
	if err := engine.Insert(elem); err != nil {
		component := FailedToAdd()
		err = component.Render(r.Context(), w)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Fatalf("Error rendering in HelloWebHandler: %e", err)
		}

		return
	}

	component := Add(elem)
	err = component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error rendering in HelloWebHandler: %e", err)
	}
}

func OnChangeSearchHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	searchedElem := r.FormValue("elem")
	fmt.Printf("Zmiana! %s\n", searchedElem)

	matched := engine.Suggestions(searchedElem, 1)
	fmt.Println(matched)
	component := Filter(matched)

	err = component.Render(r.Context(), w)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error rendering in HelloWebHandler: %e", err)
	}

}

func OnSubmitSearchHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	searchedElem := r.FormValue("elem")

	suggestions := engine.Match(searchedElem)

	component := Filter(suggestions)

	err = component.Render(r.Context(), w)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error rendering in HelloWebHandler: %e", err)
	}
}
