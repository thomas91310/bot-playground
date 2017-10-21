package main

import (
	"github.com/kataras/iris"
	"github.com/thomas91310/bot-playground/routes"
)

func main() {
	app := iris.New()

	app.Controller("/ping", new(routes.PingRoute))
	app.Controller("/webhook", new(routes.WebhookRoute))

	app.Run(iris.Addr("127.0.1:5000"))
}
