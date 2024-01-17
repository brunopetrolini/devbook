package router

import (
	"devbook/src/application/router/routes"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	r := mux.NewRouter()
	return routes.Configure(r)
}
