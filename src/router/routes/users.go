package routes

import (
	"api/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URI:      "/users",
		Method:   http.MethodPost,
		Function: controllers.UserCreate,
		Auth:     false,
	},

	{
		URI:      "/users",
		Method:   http.MethodGet,
		Function: controllers.UsersSearch,
		Auth:     true,
	},

	{
		URI:      "/users/{userId}",
		Method:   http.MethodGet,
		Function: controllers.UserSearch,
		Auth:     false,
	},

	{
		URI:      "/users/{userId}",
		Method:   http.MethodPut,
		Function: controllers.UserUpdate,
		Auth:     false,
	},

	{
		URI:      "/users/{userId}",
		Method:   http.MethodDelete,
		Function: controllers.UserDelete,
		Auth:     false,
	},
}
