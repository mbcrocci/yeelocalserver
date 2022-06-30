package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/lights", app.listLightsHandler)
	router.HandlerFunc(http.MethodGet, "/lights/:id/toggle", app.toggleLightHandler)
	router.HandlerFunc(http.MethodPost, "/lights/:id/command", app.commandLightHandler)

	return router
}
