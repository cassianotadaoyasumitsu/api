package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Rotas da API
type Route struct {
	URI      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
	Auth     bool
}

// Coloca rotas no router
func Config(r *mux.Router) *mux.Router {
	routes := usersRoutes
	routes = append(routes, LoginRoute)

	for _, route := range routes {
		if route.Auth {
			r.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Auth(route.Function)),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}
	return r
}
