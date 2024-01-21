package routes

import "net/http"

var routes = []Route{
	{
		Uri:    "/users",
		Method: http.MethodPost,
		Function: func(w http.ResponseWriter, r *http.Request) {

		},

		AuthRequire: false,
	},

	{
		Uri:    "/users",
		Method: http.MethodGet,
		Function: func(w http.ResponseWriter, r *http.Request) {

		},

		AuthRequire: false,
	},

	{
		Uri:    "/users/{userId}",
		Method: http.MethodGet,
		Function: func(w http.ResponseWriter, r *http.Request) {

		},

		AuthRequire: false,
	},

	{
		Uri:    "/users/{userId}",
		Method: http.MethodPut,
		Function: func(w http.ResponseWriter, r *http.Request) {

		},

		AuthRequire: false,
	},

	{
		Uri:    "/users/{userId}",
		Method: http.MethodDelete,
		Function: func(w http.ResponseWriter, r *http.Request) {

		},

		AuthRequire: false,
	},
}
