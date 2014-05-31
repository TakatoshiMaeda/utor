package main

import (
	"./utor"
)

func main() {
	server := utor.New()
	server.Get("/", func() string {
		return "hello world"
	})
	server.Run()
}
