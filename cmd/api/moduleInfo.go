package main

import (
	"fmt"
	"net/http"
)

func (app *application) createModuleInfoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "craete a new module_info")
}

func (app *application) showModuleInfoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDPAram(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show the details of module_info %d\n", id)
}
