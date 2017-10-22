package routes

import "fmt"
import "github.com/kataras/iris/mvc"

//FBWebhookRoute .
type FBWebhookRoute struct {
	mvc.Controller
}

//Get Ping returns Pong
func (fbWR *FBWebhookRoute) Get() string {
	fmt.Println(fbWR.Ctx)
	return "yo"
}
