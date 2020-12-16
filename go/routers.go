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
		"SignIn": "/blog/user/signin",
		"AddArticle": "/blog/user/addArticle",
		"GetArticleById": "/blog/user/article/{id}",
		"DeleteArticleById" : "/blog/user/deleteArticle/{id}",
		"CreateComment": "/blog/user/article/{id}/comments",
		"GetCommentsOfArticle": "/blog/user/article/{id}/comments",
	}
	JsonResponse(response, w, http.StatusOK)
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/blog/",
		Index,
	},

	Route{
		"DeleteArticleById",
		strings.ToUpper("Get"),
		"/blog/user/deleteArticle/{id}",
		DeleteArticleById,
	},

	Route{
		"GetArticleById",
		strings.ToUpper("Get"),
		"/blog/user/article/{id}",
		GetArticleById,
	},

	Route{
		"GetArticles",
		strings.ToUpper("Get"),
		"/blog/user/articles",
		GetArticles,
	},

	Route{
		"GetCommentsOfArticle",
		strings.ToUpper("Get"),
		"/blog/user/article/{id}/comments",
		GetCommentsOfArticle,
	},

	Route{
		"CreateComment",
		strings.ToUpper("Post"),
		"/blog/user/article/{id}/comments",
		CreateComment,
	},

	Route{
		"SignIn",
		strings.ToUpper("Get"),
		"/blog/user/signin",
		SignIn,
	},
	Route{
		Name: "AddArticle",
		Method: strings.ToUpper("Post"),
		Pattern: "/blog/user/addArticle",
		HandlerFunc: AddArticle,
	},
}
