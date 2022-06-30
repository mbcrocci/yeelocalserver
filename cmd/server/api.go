package main

import (
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mbcrocci/yeelocalsrv/internal/data"
)

func (app *application) listLightsHandler(w http.ResponseWriter, r *http.Request) {
	lights := app.models.Lights.List()

	err := app.writeJSON(w, http.StatusOK, envelope{"lights": lights}, nil)
	if err != nil {
		app.servorErrorResponse(w, r, err)
	}
}

func (app *application) toggleLightHandler(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id := params.ByName("id")
	if id == "" {
		app.badRequestResponse(w, r, errors.New("missing id"))
		return
	}

	light, err := app.models.Lights.Find(id)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	cmd := data.NewCommand(7, "toggle", make([]interface{}, 0))
	err = app.discover.SendCommand(light, cmd)
	if err != nil {
		app.servorErrorResponse(w, r, err)
		return
	}
}

func (app *application) commandLightHandler(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id := params.ByName("id")
	if id == "" {
		app.badRequestResponse(w, r, errors.New("missing id"))
		return
	}

	light, err := app.models.Lights.Find(id)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	var cmd data.Command
	err = app.readJSON(w, r, &cmd)
	if err != nil {
		app.servorErrorResponse(w, r, err)
		return
	}

	if !light.Supports(cmd.Method) {
		app.badRequestResponse(w, r, errors.New("light doesn't support method"))
		return
	}

	err = app.discover.SendCommand(light, &cmd)
	if err != nil {
		app.servorErrorResponse(w, r, err)
		return
	}
}
