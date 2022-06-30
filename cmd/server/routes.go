package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/", app.listLightsHandler)
	router.HandlerFunc(http.MethodGet, "/:id/toggle", app.toggleLightHandler)
	router.HandlerFunc(http.MethodPost, "/:id/command", app.commandLightHandler)

	return router
}
