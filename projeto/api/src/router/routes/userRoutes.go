package routes

import (
	"api/src/contollers"
	"net/http"
)

var usrRoutes = []Route{
	{
		Uri:         "/users",
		Method:      http.MethodPost,
		Function:    contollers.CreateUser,
		AuthRequire: false,
	},

	{
		Uri:         "/users",
		Method:      http.MethodGet,
		Function:    contollers.SearchUsers,
		AuthRequire: false,
	},

	{
		Uri:         "/users/{userId}",
		Method:      http.MethodGet,
		Function:    contollers.SearchUser,
		AuthRequire: false,
	},

	{
		Uri:         "/users/{userId}",
		Method:      http.MethodPut,
		Function:    contollers.UpdateUser,
		AuthRequire: false,
	},

	{
		Uri:         "/users/{userId}",
		Method:      http.MethodDelete,
		Function:    contollers.DeleteUser,
		AuthRequire: false,
	},
}
