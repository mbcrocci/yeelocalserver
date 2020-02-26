package handlers

import (
	"github.com/gofiber/fiber"
	"github.com/mbcrocci/yeelocalsrv/services"
)

type LightsHandler struct {
	repo *services.LightStore
}

func NewLightsHandler(repo *services.LightStore) *LightsHandler {
	return &LightsHandler{
		repo: repo,
	}
}

func (lh *LightsHandler) Setup(root string, app *fiber.Application) {
	app.Get(root+"/", func(c *fiber.Ctx) {
		lights := lh.repo.Lights()

		c.JSON(lights)
	})
}
