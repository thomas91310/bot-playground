package models

import (
	"fmt"
	"os"
)

//FBMessageOutput is a struct to help create a message and send it to Facebook
type FBMessageOutput struct {
	//Message .
	Message FBMessageText `json:"message"`
	//Recipient .
	Recipient FBRecipient `json:"recipient"`
}

//FBRecipient contains the id of the sender that we need to reply to
type FBRecipient struct {
	ID string `json:"id"`
}

//FBMessageText contains the content of the message that we will reply with
type FBMessageText struct {
	Text string `json:"text"`
}

//GetURL returns the URL where the message needs to be sent to
func (fbMO *FBMessageOutput) GetURL() string {
	return fmt.Sprintf("https://graph.facebook.com/v2.6/me/messages?access_token=%v", os.Getenv("FB_PAGE_ACCESS_TOKEN"))
}
