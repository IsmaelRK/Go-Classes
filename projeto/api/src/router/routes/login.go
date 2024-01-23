package routes

import "net/http"

var LoginRoute = Route{

	Uri:         "/login",
	Method:      http.MethodPost,
	Function:    func(w http.ResponseWriter, r *http.Request) {},
	AuthRequire: false,
}
