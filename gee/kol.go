package main

import (
	"fmt"

	"github.com/ryouaki/koa"
	"github.com/ryouaki/koa/example/plugin"
)

func main() {
	app := koa.New()

	app.Use(plugin.Duration)
	app.Use("/", func(ctx *koa.Context, next koa.Next) {
		fmt.Println("1")
		next()
		fmt.Println("3")
	})

	app.Get("/test", func(ctx *koa.Context, next koa.Next) {
		fmt.Println("1")
		next()
		fmt.Println("3")
	}, func(ctx *koa.Context, next koa.Next) {
		fmt.Println("2")
		// ctx.SetData("test", ctx.Query["c"][0])
		next()
		fmt.Println("5")
	}, func(ctx *koa.Context, next koa.Next) {
		// ctx.Write([]byte(ctx.GetData("test").(string)))
		fmt.Println("4")
	})

	err := app.Run(8080)
	if err != nil {
		fmt.Println(err)
	}
}
