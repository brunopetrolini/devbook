package routes

import (
	"devbook/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URI:         "/users",
		Method:      http.MethodPost,
		Function:    controllers.CreateUser,
		RequireAuth: false,
	},
	{
		URI:         "/users",
		Method:      http.MethodGet,
		Function:    controllers.GetUsers,
		RequireAuth: false,
	},
	{
		URI:         "/users/{id}",
		Method:      http.MethodGet,
		Function:    controllers.FindUser,
		RequireAuth: false,
	},
	{
		URI:         "/users/{id}",
		Method:      http.MethodPut,
		Function:    controllers.UpdateUser,
		RequireAuth: false,
	},
	{
		URI:         "/users/{id}",
		Method:      http.MethodDelete,
		Function:    controllers.DeleteUser,
		RequireAuth: false,
	},
}
