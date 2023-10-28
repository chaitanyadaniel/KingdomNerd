package routes

import "net/http"

// Route represents a single route
type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}
