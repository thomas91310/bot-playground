package routes

import (
	"github.com/kataras/iris/mvc"
	"github.com/thomas91310/bot-playground/models"
)

type PingRoute struct {
	mvc.Controller
}

//Get Ping returns Pong
func (pr *PingRoute) Get() string {
	resp := models.MakeResp(200, models.Ping{
		Message: "PONG",
	})
	return resp
}
