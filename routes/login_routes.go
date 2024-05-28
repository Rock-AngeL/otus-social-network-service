package routes

import (
	"otus-social-network-service_gen_swagger/handlers"
	"otus-social-network-service_gen_swagger/repository"
	"strings"
)

func LoginRoutes(rm *repository.RepositoryManager) []Route {
	routes := make([]Route, 0)
	routes = append(routes, Route{
		Name:        "LoginPost",
		Method:      strings.ToUpper("Post"),
		Pattern:     "/login",
		HandlerFunc: handlers.NewLoginHandler(rm).LoginPost,
	})
	return routes
}
