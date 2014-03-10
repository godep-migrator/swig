package intake

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	kafka "github.com/Shopify/sarama"
	"log"
	"net/http"
)

func Create(producer *kafka.Producer, req *http.Request, res http.ResponseWriter, params martini.Params, logger *log.Logger, r render.Render) {
	logger.Println(params)
  logger.Println(producer)

	err := req.ParseForm()
	if err != nil {
		logger.Println(err)
	}

	logger.Println(req.Form)

	r.JSON(200, map[string]bool{"created": true})
}
