package utor

import (
	"log"
	"net/http"
)

type Utor struct {
	router Router
}

func New() *Utor {
	server := &Utor{router: &Router{}}
	return server
}

func (u *Utor) ServHTTP(res http.ResponseWriter, req *http.Request) {
	route, err := u.router.Find(req)

	if err != nil {
		http.Error(res, "NotfoundRoute", 404)
		return
	}

	route.Run(res, req)
}

func (u *Utor) Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("listning port on ", port)

	http.ServHTTP("localhost:"+port, u)
}

func (u *Utor) Get(path string, fn RouteFunc) {
	u.router.Add("GET", path, fn)
}
