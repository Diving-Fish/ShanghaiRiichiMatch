package main

import (
	"github.com/kataras/iris"
)
import "match/controller"

func main() {
	app := iris.Default()
	crs := func(ctx iris.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Origin,Content-Type,Authorization")
		if ctx.Method() == "OPTIONS" {
			return
		}
		ctx.Next()
	}
	app.Use(crs)
	admin := app.Party("/api/admin", controller.AdminHandler).AllowMethods(iris.MethodOptions)
	{
		admin.Get("/apply", controller.ApplyNewPlayer)
	}
	player := app.Party("/api/player", controller.AuthHandler).AllowMethods(iris.MethodOptions)
	{
		player.Get("/status", controller.Status)
		player.Post("/change_pwd", controller.ChangePwd)
		player.Post("/bind", controller.Bind)
	}
	public := app.Party("/api/public").AllowMethods(iris.MethodOptions)
	{
		public.Post("/login", controller.Login)
		public.Get("/search_player", controller.SearchPlayerById)
	}
	_ = app.Run(iris.Addr(":8088"))
}
