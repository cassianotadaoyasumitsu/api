package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Routes
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Config(r)
}
