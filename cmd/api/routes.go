package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthCheckerHandler)
	router.HandlerFunc(http.MethodGet, "/v1/moduleinfo/:id", app.showModuleInfoHandler)
	router.HandlerFunc(http.MethodPost, "/v1/moduleinfo", app.createModuleInfoHandler)

	return router
}