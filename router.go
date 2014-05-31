package utor

import (
	"net/http"
)

type Router struct {
	routes []*Route
}

func (r *Router) Add(method string, path string, fn RouteFunc) {
	route := &Route{method: method, path: path, call: fn}
	r.routes = append(r.routes, route)
}

func (r *Router) Find(request *http.Request) {
	for _, route := range r.routes {
		if route.Match(request) {
			return route
		}
	}
	return nil
}
