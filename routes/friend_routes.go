package routes

import (
	"otus-social-network-service_gen_swagger/handlers"
	"otus-social-network-service_gen_swagger/repository"
	"strings"
)

func FriendRoutes(rm *repository.RepositoryManager) []Route {
	routes := make([]Route, 0)
	routes = append(routes,
		Route{
			Name:        "FriendDeleteUserIdPut",
			Method:      strings.ToUpper("Put"),
			Pattern:     "/friend/delete/{user_id}",
			HandlerFunc: handlers.NewFriendHandler(rm).FriendDeleteUserIdPut,
		},
		Route{
			Name:        "FriendSetUserIdPut",
			Method:      strings.ToUpper("Put"),
			Pattern:     "/friend/set/{user_id}",
			HandlerFunc: handlers.NewFriendHandler(rm).FriendSetUserIdPut,
		})
	return routes
}
