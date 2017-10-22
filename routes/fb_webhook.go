package routes

import (
	"os"

	"github.com/kataras/iris/mvc"
	"github.com/thomas91310/bot-playground/models"
)

//FBWebhookRoute .
type FBWebhookRoute struct {
	mvc.Controller
}

//Get Ping returns Pong
func (fbWR *FBWebhookRoute) Get() string {
	qs := fbWR.Ctx.Request().URL.Query()
	token, exists := qs["hub.verify_token"]
	if !exists {
		return models.MakeBadResp(400, "Invalid request")
	}
	if token[0] != os.Getenv("FB_APP_TOKEN") {
		return models.MakeBadResp(400, "Invalid hub.verify_token")
	}
	expectedFromFB, exists := qs["hub.challenge"]
	if !exists {
		return models.MakeBadResp(400, "Invalid request")
	}
	return expectedFromFB[0]
}
