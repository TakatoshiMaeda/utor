package utor

import (
	"net/http"
)

type RouteFunc interface{}

type Route struct {
	method string
	path   string
	call   RouteFunc
}

func (r *Route) Match(request *http.Request) bool {
	if request.Method == r.method && request.URL.Path == r.path {
		return true
	}
	return false
}

func (r *Route) Run(response http.ResponseWriter, request *http.Request) {
	response.Write(r.call())
}
