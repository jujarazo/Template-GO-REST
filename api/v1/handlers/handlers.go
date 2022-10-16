package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/manga", getMangaList)

	return r
}

func getMangaList(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte{})
}
