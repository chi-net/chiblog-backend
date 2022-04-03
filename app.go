/*
 * @Author: chihuo2104
 * @Date: 2022-04-03 16:27:09
 * @LastEditTime: 2022-04-03 16:35:29
 * @LastEditors: chihuo2104
 * @Description: go app file
 * @FilePath: \chiblog-backend\app.go
 * Powered by chihuo2104<chihuo2104@moekonnyaku.com>.All right reserved ©2018-now.
 */
package main

import (
	"github.com/chihuo2104/chiblog-backend/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// posts comments settings 均用sw存储
	app.Get("/healthcheck", routes.Check)
	app.Get("/posts", routes.GetPosts)
	app.Get("/comments", routes.GetComments)
	app.Get("/settings", routes.GetSettings)

	// Post
	app.Post("/posts/create", routes.CreatePosts)
	app.Post("/comments/create", routes.CreateComments)
	app.Post("/settings/update", routes.UpdateSettings)
	app.Post("/login", routes.Login)

	// 404兜底
	app.All("*", routes.F0f)
	// Listen
	app.Listen(":3000")
}
