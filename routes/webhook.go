package routes

import (
	"github.com/kataras/iris/mvc"
	"github.com/thomas91310/bot-playground/models"
)

//WebhookRoute .
type WebhookRoute struct {
	mvc.Controller
}

//Get a Facebook Webhook
func (*WebhookRoute) Get() string {
	resp := models.MakeResp(200, models.WebHook{
		Message: "YO",
	})
	return resp
}
