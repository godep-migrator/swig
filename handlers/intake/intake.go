package intake

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	//"github.com/Shopify/sarama"
	"log"
	"net/http"
)

func Create(req *http.Request, res http.ResponseWriter, params martini.Params, logger *log.Logger, r render.Render) {
	logger.Println(params)

	err := req.ParseForm()
	if err != nil {
		logger.Panic(err)
	}

	logger.Println(req.Form)

	r.JSON(200, map[string]bool{"created": true})
}
