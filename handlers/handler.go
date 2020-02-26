package handlers

import "github.com/gofiber/fiber"

type Handler interface {
	Setup(root string, app *fiber.Application)
}
