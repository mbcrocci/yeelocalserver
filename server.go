package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber"
	"github.com/mbcrocci/yeelocalsrv/entities"
	"github.com/mbcrocci/yeelocalsrv/handlers"
	"github.com/mbcrocci/yeelocalsrv/services"
)

type Server interface {
	Init()
	Start()
	Shutdown()
}

type LightsServer struct {
	app           *fiber.App
	port          int
	lightsHandler handlers.Handler
	discover      *services.DiscoverService
	lightChannel  chan string
	repo          *services.LightStore
	logger        *log.Logger
}

func (s *LightsServer) Init() {
	s.logger = log.New(os.Stdout, "lights ", log.LstdFlags)

	s.repo = &services.LightStore{}
	s.repo.Init()

	s.discover = services.NewDiscoverService()
	s.discover.Init()

	s.app = fiber.New()
	s.lightsHandler = handlers.NewLightsHandler(s.repo, s.discover, s.logger)

	s.lightsHandler.Setup("/lights", s.app)

}

func (s *LightsServer) Start() {
	s.ListenLights()
	s.app.Listen(s.port)
}

func (s *LightsServer) Shutdown() error {
	log.Println("Shuting down server...")
	err := s.app.Shutdown()
	if err != nil {
		return err
	}

	err = s.discover.Shutdown()
	if err != nil {
		return err
	}

	return nil
}

func (s *LightsServer) ListenLights() {
	lc := make(chan *entities.Light)
	ec := make(chan error)
	s.discover.Start(lc, ec)

	go func() {
		for {
			select {
			case light := <-lc:
				s.repo.Add(light)

			case err := <-ec:
				s.logger.Println(err)
			}
		}
	}()
}
