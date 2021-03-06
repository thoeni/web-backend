package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/api",
		Index,
	},
	Route{
		"Get page",
		"GET",
		"/api/page/{pageId}",
		PageResource,
	},
	Route{
		"Get page",
		"GET",
		"/api/contacts",
		ContactsResource,
	},
	Route{
		"Get testimonials",
		"GET",
		"/api/testimonial",
		AllTestimonialResource,
	},
	Route{
		"Get specific testimonial",
		"GET",
		"/api/testimonial/{name}",
		TestimonialResource,
	},
	Route{
		"Get projects",
		"GET",
		"/api/project",
		AllProjectResource,
	},
	Route{
		"Get specific project",
		"GET",
		"/api/project/{name}",
		ProjectResource,
	},
}
