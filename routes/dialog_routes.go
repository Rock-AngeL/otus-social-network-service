package routes

import (
	"otus-social-network-service_gen_swagger/handlers"
	"otus-social-network-service_gen_swagger/repository"
	"strings"
)

func DialogRoutes(rm *repository.RepositoryManager) []Route {
	routes := make([]Route, 0)
	routes = append(routes,
		Route{
			Name:        "DialogUserIdListGet",
			Method:      strings.ToUpper("Get"),
			Pattern:     "/dialog/{user_id}/list",
			HandlerFunc: handlers.NewDialogHandler(rm).DialogUserIdListGet,
		},
		Route{
			Name:        "DialogUserIdSendPost",
			Method:      strings.ToUpper("Post"),
			Pattern:     "/dialog/{user_id}/send",
			HandlerFunc: handlers.NewDialogHandler(rm).DialogUserIdSendPost,
		})
	return routes
}
