package routes

import (
	"api/src/controllers"
	"net/http"
)

var PostsRoutes = []Route{
	{
		URI:      "/posts",
		Method:   http.MethodPost,
		Function: controllers.PostCreate,
		Auth:     true,
	},

	{
		URI:      "/posts",
		Method:   http.MethodGet,
		Function: controllers.PostsSearch,
		Auth:     true,
	},

	{
		URI:      "/posts/{postId}",
		Method:   http.MethodGet,
		Function: controllers.PostSearch,
		Auth:     true,
	},

	{
		URI:      "/posts/{postId}",
		Method:   http.MethodPut,
		Function: controllers.PostUpdate,
		Auth:     true,
	},

	{
		URI:      "/posts/{postId}",
		Method:   http.MethodDelete,
		Function: controllers.PostDelete,
		Auth:     true,
	},

	{
		URI:      "/users/{userId}/posts",
		Method:   http.MethodGet,
		Function: controllers.PostsByUser,
		Auth:     true,
	},

	{
		URI:      "/posts/{postId}/likes",
		Method:   http.MethodPost,
		Function: controllers.LikePost,
		Auth:     true,
	},
}
