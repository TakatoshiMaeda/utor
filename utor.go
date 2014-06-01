package utor

import (
	"log"
	"net/http"
	"os"
)

type Utor struct {
	http.Handler
	router Router
}

func New() *Utor {
	server := &Utor{}
	return server
}

func (u *Utor) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	route := u.router.Find(req)

	if route == nil {
		http.Error(res, "Notfound", 404)
		return
	}

	route.Run(res, req)
}

func (u *Utor) Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("listning port on %s", port)

	http.ListenAndServe("localhost:"+port, u)
}

func (u *Utor) Get(path string, fn RouteFunc) {
	u.router.Add("GET", path, fn)
}

func (u *Utor) Post(path string, fn RouteFunc) {
	u.router.Add("POST", path, fn)
}

func (u *Utor) Put(path string, fn RouteFunc) {
	u.router.Add("PUT", path, fn)
}

func (u *Utor) Delete(path string, fn RouteFunc) {
	u.router.Add("DELETE", path, fn)
}

func (u *Utor) Option(path string, fn RouteFunc) {
	u.router.Add("OPTION", path, fn)
}

func (u *Utor) Head(path string, fn RouteFunc) {
	u.router.Add("HEAD", path, fn)
}
