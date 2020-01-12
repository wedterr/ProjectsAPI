package controller

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Route contains route related properties.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	//router.Use(commonMiddleware)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Projects API!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"GetAuthors",
		"GET",
		"/authors",
		GetAuthors,
	},
	Route{
		"GetAuthor",
		"GET",
		"/authors/{id}",
		GetAuthor,
	},
	Route{
		"CreateAuthor",
		"POST",
		"/authors",
		CreateAuthor,
	},
	Route{
		"RemoveAuthor",
		"DELETE",
		"/authors/{id}",
		DeleteAuthor,
	},
	Route{
		"UdateAuthor",
		"PUT",
		"/authors/{id}",
		UpdateAuthor,
	},

	// projects
	// Bad logic, need to make projects as a separate entity not embeded in authors
	Route{
		"GetProjects",
		"GET",
		"/authors/{authorid}/projects",
		GetProjects,
	},
	Route{
		"GetProject",
		"GET",
		"/authors/{authorid}/projects/{projectid}",
		GetProject,
	},
	/*
		Route{
			"CreateProject",
			"POST",
			"/authors/{authorid}/projects",
			CreateProject,
		},
			Route{
				"RemoveProject",
				"DELETE",
				"/authors/{authorid}/projects/{projectid}",
				DeleteProject,
			},
			Route{
				"UdateProject",
				"PUT",
				"/authors/{authorid}/projects/{projectid}",
				UpdateProject,
			},
	*/

	Route{
		"GetToken",
		"GET",
		"/token",
		GetToken,
	},
}
