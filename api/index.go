package handler

import (
	"fmt"
	"github.com/tbxark/g4vercel"
	"net/http"
)


func Handler(w http.ResponseWriter, r *http.Request) {
	server := gee.New()
	server.Use(gee.Recovery(func(err interface{}, c *gee.Context) bool {
		message := fmt.Sprintf("%s", err)
		c.JSON(500, gee.H{
			"error": message,
		})
		return true
	}))
	server.GET("/", func(context *gee.Context) {
		context.JSON(200, gee.H{
			"status": "OK",
		})
	})
	server.GET("/hello", func(context *gee.Context) {
		name := context.Query("name")
		if name == "" {
			context.JSON(400, gee.H{
				"error": "name not found",
			})
		} else {
			context.JSON(200, gee.H{
				"data": fmt.Sprintf("Hello %s!", name),
			})
		}
	})
	server.GET("/user/:id", func(context *gee.Context) {
		context.JSON(400, gee.H{
			"data": gee.H{
				"id": context.Param("id"),
			},
		})
	})
	server.GET("/long/long/long/path/*test", func(context *gee.Context) {
		context.JSON(200, gee.H{
			"data": gee.H{
				"url": context.Path,
			},
		})
	})
	server.Handle(w, r)
}
