package routes

import (
	"github.com/codegangsta/martini"
	"github.com/jeffchao/swig/handlers/intake"
	"github.com/jeffchao/swig/handlers/root"
)

func Route(server *martini.ClassicMartini) {
	server.Get("/", root.Index)
	server.Post("/", intake.Create)
}
