package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthCheckerHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "status: available\n")
	fmt.Fprintf(writer, "environment: %s\n", app.config.env)
	fmt.Fprintf(writer, "version: %s\n", version)
}
