package handlers

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/mbcrocci/yeelocalsrv/entities"
	"github.com/mbcrocci/yeelocalsrv/services"
)

type LightsHandler struct {
	repo     *services.LightStore
	discover *services.DiscoverService
}

func NewLightsHandler(repo *services.LightStore, ds *services.DiscoverService) *LightsHandler {
	return &LightsHandler{
		repo:     repo,
		discover: ds,
	}
}

func (lh *LightsHandler) Setup(root string, app *fiber.App) {
	app.Get(root+"/", func(c *fiber.Ctx) {
		lights := lh.repo.Lights()

		c.JSON(lights)
	})

	app.Get(root+"/toggle", func(c *fiber.Ctx) {
		lights := lh.repo.Lights()
		cmd := entities.NewCommand(7, "toggle", make([]string, 0))

		for _, light := range lights {
			err := lh.discover.SendCommand(light, cmd)
			if err != nil {
				fmt.Println(err)
				c.SendStatus(500)
				return
			}
		}

		c.SendStatus(200)
	})
}
