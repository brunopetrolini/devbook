package routes

import (
	"devbook/src/application/controllers"
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
		RequireAuth: true,
	},
	{
		URI:         "/users/{id}",
		Method:      http.MethodGet,
		Function:    controllers.FindUser,
		RequireAuth: true,
	},
	{
		URI:         "/users/{id}",
		Method:      http.MethodPut,
		Function:    controllers.UpdateUser,
		RequireAuth: true,
	},
	{
		URI:         "/users/{id}",
		Method:      http.MethodDelete,
		Function:    controllers.DeleteUser,
		RequireAuth: true,
	},
	{
		URI:         "/users/{id}/follow",
		Method:      http.MethodPost,
		Function:    controllers.FollowUser,
		RequireAuth: true,
	},
}
