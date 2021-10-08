package main

import (
	"app/controller"
	"app/database"
	"app/environment"
	"app/service"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := iris.New()
	app.Get("/ping", pong).Describe("healthcheck")

	mvc.Configure(app.Party("/greet"), setup)

	// http://localhost:8080/greet?name=kataras
	app.Listen(":8080", iris.WithLogLevel("debug"))
}

func pong(ctx iris.Context) {
	ctx.WriteString("pong")
}

func setup(app *mvc.Application) {
	// Register Dependencies.
	app.Register(
		environment.DEV,         // DEV, PROD
		database.NewDB,          // sqlite, mysql
		service.NewGreetService, // greeterWithLogging, greeter
	)

	// Register Controllers.
	app.Handle(new(controller.GreetController))
}
