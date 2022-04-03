/*
 * @Author: chihuo2104
 * @Date: 2022-04-03 16:27:09
 * @LastEditTime: 2022-04-03 16:35:56
 * @LastEditors: chihuo2104
 * @Description: routes-index
 * @FilePath: \chiblog-backend\routes\index.go
 * Powered by chihuo2104<chihuo2104@moekonnyaku.com>.All right reserved ©2018-now.
 */
package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	return c.SendString("What a stupid taromaru it is!")
}
