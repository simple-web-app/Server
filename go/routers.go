/*
 * simple web blog
 *
 * Simple Web Blog
 *
 * API version: 1.0.0
 * Contact: nanzh@mail2.sysu.edu.cn
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"net/http"
	"strings"

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

func Index(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"SignIn": "/api/user/signin",
		"AddArticle": "/api/article",
		"GetArticleById": "/api/article/{id}",
		"DeleteArticleById" : "/api/article/{id}",
		"CreateComment": "/api/article/{id}/comments",
		"GetCommentsOfArticle": "/api/article/{id}/comments",
		"GetTags": "/api/tag",
		"GetTagById": "/api/tag/{id}",
		"AddTag": "/api/tag",
	}
	JsonResponse(response, w, http.StatusOK)
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/api/",
		Index,
	},

	Route{
		"DeleteArticleById",
		strings.ToUpper("Delete"),
		"/api/article/{id}",
		DeleteArticleById,
	},

	Route{
		"GetArticleById",
		strings.ToUpper("Get"),
		"/api/article/{id}",
		GetArticleById,
	},

	Route{
		"GetArticles",
		strings.ToUpper("Get"),
		"/api/article",
		GetArticles,
	},

	Route{
		"GetCommentsOfArticle",
		strings.ToUpper("Get"),
		"/api/article/{id}/comments",
		GetCommentsOfArticle,
	},

	Route{
		"CreateComment",
		strings.ToUpper("Post"),
		"/api/article/{id}/comments",
		CreateComment,
	},

	Route{
		"SignIn",
		strings.ToUpper("Get"),
		"/api/user/signin",
		SignIn,
	},
	Route{
		Name: "AddArticle",
		Method: strings.ToUpper("Post"),
		Pattern: "/api/article",
		HandlerFunc: AddArticle,
	},
	Route{
		Name: "GetTags",
		Method: strings.ToUpper("Get"),
		Pattern: "/api/tag",
		HandlerFunc: GetTags,
	},
	Route{
		Name: "GetTagById",
		Method: strings.ToUpper("Get"),
		Pattern: "/api/tag/{id}",
		HandlerFunc: GetTagById,
	},
	Route{
		Name: "AddTag",
		Method: strings.ToUpper("Post"),
		Pattern: "/api/tag",
		HandlerFunc: AddTag,
	},
}
