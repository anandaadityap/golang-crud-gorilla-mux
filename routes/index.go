package routes

import "github.com/gorilla/mux"

func RoutesIndex(r *mux.Router) {
	api := r.PathPrefix("/api").Subrouter()

	AuthorRoutes(api)
	BooksRoutes(api)
}
