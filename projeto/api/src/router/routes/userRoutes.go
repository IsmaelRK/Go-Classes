package routes

import (
	"api/src/controllers"
	"net/http"
)

var usrRoutes = []Route{
	{
		Uri:         "/users",
		Method:      http.MethodPost,
		Function:    controllers.CreateUser,
		AuthRequire: false,
	},

	{
		Uri:         "/users",
		Method:      http.MethodGet,
		Function:    controllers.SearchUsers,
		AuthRequire: false,
	},

	{
		Uri:         "/users/{userId}",
		Method:      http.MethodGet,
		Function:    controllers.SearchUser,
		AuthRequire: false,
	},

	{
		Uri:         "/users/{userId}",
		Method:      http.MethodPut,
		Function:    controllers.UpdateUser,
		AuthRequire: false,
	},

	{
		Uri:         "/users/{userId}",
		Method:      http.MethodDelete,
		Function:    controllers.DeleteUser,
		AuthRequire: false,
	},
}
