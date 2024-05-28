package routes

import (
	"otus-social-network-service_gen_swagger/handlers"
	"otus-social-network-service_gen_swagger/repository"
	"strings"
)

func PostRoutes(rm *repository.RepositoryManager) []Route {
	routes := make([]Route, 0)
	routes = append(routes,
		Route{
			Name:        "PostCreatePost",
			Method:      strings.ToUpper("Post"),
			Pattern:     "/post/create",
			HandlerFunc: handlers.NewPostHandler(rm).PostCreatePost,
		},
		Route{
			Name:        "PostDeleteIdPut",
			Method:      strings.ToUpper("Put"),
			Pattern:     "/post/delete/{id}",
			HandlerFunc: handlers.NewPostHandler(rm).PostDeleteIdPut,
		},
		Route{
			Name:        "PostFeedGet",
			Method:      strings.ToUpper("Get"),
			Pattern:     "/post/feed",
			HandlerFunc: handlers.NewPostHandler(rm).PostFeedGet,
		},
		Route{
			Name:        "PostGetIdGet",
			Method:      strings.ToUpper("Get"),
			Pattern:     "/post/get/{id}",
			HandlerFunc: handlers.NewPostHandler(rm).PostGetIdGet,
		},
		Route{
			Name:        "PostUpdatePut",
			Method:      strings.ToUpper("Put"),
			Pattern:     "/post/update",
			HandlerFunc: handlers.NewPostHandler(rm).PostUpdatePut,
		},
	)
	return routes
}
