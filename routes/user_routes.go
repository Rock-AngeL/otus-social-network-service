package routes

import (
	"otus-social-network-service_gen_swagger/handlers"
	"otus-social-network-service_gen_swagger/repository"
	"strings"
)

func UserRoutes(rm *repository.RepositoryManager) []Route {
	routes := make([]Route, 0)
	routes = append(routes,
		Route{
			Name:        "UserGetIdGet",
			Method:      strings.ToUpper("Get"),
			Pattern:     "/user/get/{id}",
			HandlerFunc: handlers.NewUserHandler(rm).UserGetIdGet,
		},
		Route{
			Name:        "UserRegisterPost",
			Method:      strings.ToUpper("Post"),
			Pattern:     "/user/register",
			HandlerFunc: handlers.NewUserHandler(rm).UserRegisterPost,
		},
		Route{
			Name:        "UserSearchGet",
			Method:      strings.ToUpper("Get"),
			Pattern:     "/user/search",
			HandlerFunc: handlers.NewUserHandler(rm).UserSearchGet,
		},
	)
	return routes
}
