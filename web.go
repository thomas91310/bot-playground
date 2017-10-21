package main

import (
	"fmt"
	"os"

	"github.com/kataras/iris"
	"github.com/thomas91310/bot-playground/routes"
)

func main() {
	app := iris.New()

	app.Controller("/ping", new(routes.PingRoute))
	app.Controller("/webhook", new(routes.WebhookRoute))
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	address := fmt.Sprintf("0.0.0.0:%v", port)
	app.Run(iris.Addr(address))
}
