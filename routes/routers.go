/*
 * OTUS Highload Architect
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"otus-social-network-service_gen_swagger/logger"
	"otus-social-network-service_gen_swagger/repository"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter(rm *repository.RepositoryManager) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	routes := make(Routes, 0)
	routes = append(routes, LoginRoutes(rm)...)
	routes = append(routes, UserRoutes(rm)...)
	routes = append(routes, PostRoutes(rm)...)
	routes = append(routes, FriendRoutes(rm)...)
	routes = append(routes, DialogRoutes(rm)...)
	routes = append(routes, Route{
		"Index",
		"GET",
		"/",
		Index,
	})

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}