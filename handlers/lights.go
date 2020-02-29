package handlers

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber"
	"github.com/mbcrocci/yeelocalsrv/entities"
	"github.com/mbcrocci/yeelocalsrv/services"
)

type LightsHandler struct {
	repo     *services.LightStore
	discover *services.DiscoverService
	logger   *log.Logger
}

func NewLightsHandler(repo *services.LightStore, ds *services.DiscoverService, logger *log.Logger) *LightsHandler {
	return &LightsHandler{
		repo:     repo,
		discover: ds,
		logger:   logger,
	}
}

func (lh *LightsHandler) Setup(root string, app *fiber.App) {
	app.Get(root+"/", func(c *fiber.Ctx) {
		lights := lh.repo.Lights()

		c.JSON(lights)
	})

	app.Get(root+"/toggle", func(c *fiber.Ctx) {
		lights := lh.repo.Lights()
		cmd := entities.NewCommand(7, "toggle", make([]interface{}, 0))

		for _, light := range lights {
			err := lh.discover.SendCommand(light, cmd)
			if err != nil {
				lh.logger.Println(err)
				c.SendStatus(http.StatusInternalServerError)
				return
			}
		}

		c.SendStatus(http.StatusOK)
	})

	app.Post(root+"/:id/command", func(c *fiber.Ctx) {
		id := c.Params("id")
		light, err := lh.repo.Find(id)
		if err != nil {
			c.SendStatus(http.StatusNotFound)
			return
		}

		var cmd entities.Command
		if err := c.BodyParser(&cmd); err != nil {
			lh.logger.Println(err)

			c.SendStatus(http.StatusBadRequest)
			return
		}

		if !light.Supports(cmd.Method) {
			c.SendStatus(http.StatusMethodNotAllowed)
			return
		}

		err = lh.discover.SendCommand(light, &cmd)
		if err != nil {
			lh.logger.Println(err)
			c.SendStatus(http.StatusInternalServerError)
			return
		}

		c.SendStatus(http.StatusOK)
	})
}
