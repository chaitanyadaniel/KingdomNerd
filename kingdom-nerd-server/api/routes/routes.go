package routes

import (
	"server/package/handlers/home"
	"server/package/handlers/login"
	"server/package/middleware"
)

// getRoutes returns an array of routes
func GetRoutes() []Route {
	routes := []Route{
		{Path: "/home", Method: "GET", Handler: middleware.AuthenticateMiddleware()(home.HomeHandler)},
		//{Path: "/home", Method: "GET", Handler: home.HomeHandler},
		{Path: "/login", Method: "POST", Handler: login.LoginHandler},
		{Path: "/loginCheck", Method: "GET", Handler: middleware.AuthenticateMiddleware()(login.LoginCheckHandler)},
		// Add more routes here
	}

	return routes
}
