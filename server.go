package main

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/jeffchao/swig/config/kafka"
	"github.com/jeffchao/swig/config/routes"
	"net/http"
)

func main() {
	server := martini.Classic()
	routes.Route(server)

	server.Use(render.Renderer(render.Options{IndentJSON: true}))
	server.Use(kafka.Kafka())

	http.ListenAndServe(":8080", server)
	server.Run()
}
