package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthCheckerHandler)
	router.HandlerFunc(http.MethodGet, "/v1/moduleinfo/:id", app.showModuleInfoHandler)
	router.HandlerFunc(http.MethodPost, "/v1/moduleinfo", app.createModuleInfoHandler)
	router.HandlerFunc(http.MethodPut, "/v1/moduleinfo/:id", app.UpdateModuleInfoHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/moduleinfo/:id", app.DeleteModuleInfoHandler)

	return router
}
