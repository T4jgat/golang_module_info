package main

import (
	"fmt"
	"github.com/T4jgat/module_info/internal/data"
	"net/http"
	"time"
)

func (app *application) createModuleInfoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "craete a new module_info")
}

func (app *application) showModuleInfoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDPAram(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	module := data.ModuleInfo{
		ID:             id,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		ModuleName:     "",
		ModuleDuration: 0,
		ExamType:       []string{""},
		Runtime:        102,
		Version:        1,
	}

	err = app.writeJSON(w, http.StatusOK, module, nil)
	if err != nil {
		app.logger.Print(err)
		app.serverErrorResponse(w, r, err)
	}

}
