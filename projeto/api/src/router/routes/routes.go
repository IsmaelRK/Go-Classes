package routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Uri         string
	Method      string
	Function    func(http.ResponseWriter, *http.Request)
	AuthRequire bool
}

func ConfigRoutes(r *mux.Router) *mux.Router {

	routes := usrRoutes

	for _, route := range routes {

		r.HandleFunc(route.Uri, route.Function).Methods(route.Method)

	}

	return r
}
