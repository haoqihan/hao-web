package main

import (
	"hao-web/hao"
	"net/http"
)

func main() {
	r := hao.New()
	r.GET("/", func(ctx *hao.Context) {
		ctx.HTML(http.StatusOK, "<h1>hello hao</h1>")
	})
	r.GET("/hello", func(ctx *hao.Context) {
		ctx.String(http.StatusOK, "hello %s,you're at %s\n", ctx.Query("name"), ctx.Path)
	})
	r.POST("/login", func(ctx *hao.Context) {
		ctx.JSON(http.StatusOK, hao.H{
			"username": ctx.PostForm("username"),
			"password": ctx.PostForm("password"),
		})
	})
	r.Run(":9999")
}
