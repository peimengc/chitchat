package routes

import (
	"github.com/gorilla/mux"
	"github.com/peimengc/chitchat/handlers"
	"net/http"
)

type WebRoute struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type WebRoutes []WebRoute

var webRoutes = WebRoutes{
	{
		"home",
		"get",
		"/",
		handlers.Index,
	}, {
		"login",
		"get",
		"/login",
		handlers.Login,
	}, {
		"auth",
		"post",
		"/authenticate",
		handlers.Authenticate,
	}, {
		"signupAccount",
		"post",
		"/signup_account",
		handlers.SignupAccount,
	}, {
		"signup",
		"get",
		"/signup",
		handlers.Signup,
	}, {
		"logout",
		"get",
		"/logout",
		handlers.Logout,
	}, {
		"newThread",
		"get",
		"/thread/new",
		handlers.NewThread,
	}, {
		"createThread",
		"post",
		"/thread/create",
		handlers.CreateThread,
	}, {
		"readThread",
		"get",
		"/thread/read",
		handlers.ReadThread,
	}, {
		"postThread",
		"post",
		"/thread/post",
		handlers.PostThread,
	}, {
		"error",
		"get",
		"/err",
		handlers.Err,
	},
}

//加载路由
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range webRoutes {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}
