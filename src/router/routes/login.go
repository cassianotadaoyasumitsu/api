package routes

import (
	"api/src/controllers"
	"net/http"
)

var LoginRoute = Route{
	URI:      "/login",
	Method:   http.MethodPost,
	Function: controllers.Login,
	Auth:     false,
}
