package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/kataras/iris/mvc"
	"github.com/thomas91310/bot-playground/models"
)

//FBWebhookRoute handles facebook requests
type FBWebhookRoute struct {
	mvc.Controller
}

//Get FBWebhook returns what Facebook expects from the get webhook endpoint
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

//Post FBWebhook handles posts from facebook
//Messenger messages from people to the bot are coming in here
func (fbWR *FBWebhookRoute) Post() {
	FBmessageInput := new(models.FBMessageInput)

	body, err := ioutil.ReadAll(fbWR.Ctx.Request().Body)
	err = json.Unmarshal(body, FBmessageInput)
	if err != nil {
		log.Printf("Error unmarshalling message from facebook: %v. Got %v", body, err)
		return
	}

	fbWR.Respond(FBmessageInput.Entry[0].Messaging[0].Sender.ID)
}

//Respond creates a message and sends it to the sender
func (fbWR *FBWebhookRoute) Respond(senderID string) error {
	message := models.FBMessageOutput{
		Recipient: models.FBRecipient{
			ID: senderID,
		},
		Message: models.FBMessageText{
			Text: "I'm just a little something",
		},
	}

	bb := new(bytes.Buffer)
	err := json.NewEncoder(bb).Encode(message)
	if err != nil {
		return err
	}

	resp, err := http.Post(
		message.GetURL(),
		"application/json",
		bb,
	)
	if err != nil {
		return err
	}

	fmt.Println(string(bb.Bytes()))
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("resp : ", string(b))

	return nil
}
