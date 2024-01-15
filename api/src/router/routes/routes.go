package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route represents a route in the API
type Route struct {
	URI         string
	Method      string
	Function    func(http.ResponseWriter, *http.Request)
	RequireAuth bool
}

// Configure initializes the routes in the mux router
func Configure(r *mux.Router) *mux.Router {
	routes := usersRoutes
	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}
	return r
}
