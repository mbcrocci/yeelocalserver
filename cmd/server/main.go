package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mbcrocci/yeelocalsrv/internal/data"
	"github.com/mbcrocci/yeelocalsrv/internal/yeelight"
)

type config struct {
	port int
}

type application struct {
	config config
	logger *log.Logger

	models   data.Models
	discover *yeelight.DiscoverService
	//server *lights.Server
}

func (app *application) syncChanges(lc chan *data.Light, ec chan error) {
	for {
		select {
		case light := <-lc:
			app.models.Lights.Add(light)

		case err := <-ec:
			app.logger.Println(err)
		}
	}
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 3000, "Lights Server")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	app := &application{
		config:   cfg,
		logger:   logger,
		models:   data.NewModels(),
		discover: yeelight.NewDiscoverService(),
	}

	lights, errC, err := app.discover.Start()
	if err != nil {
		logger.Fatal(err)
	}
	go app.syncChanges(lights, errC)

	stv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting server on %s", cfg.port)

	err = stv.ListenAndServe()
	logger.Fatal(err)
}
