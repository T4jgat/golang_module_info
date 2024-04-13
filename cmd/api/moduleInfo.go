package main

import (
	"errors"
	"fmt"
	"github.com/T4jgat/module_info/internal/data"
	"github.com/T4jgat/module_info/internal/validator"
	"net/http"
)

func (app *application) createModuleInfoHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ModuleName     string   `json:"module_name"`
		ModuleDuration int32    `json:"module_duration"`
		ExamType       []string `json:"exam_type"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	moduleInfo := &data.ModuleInfo{
		ModuleName:     input.ModuleName,
		ModuleDuration: input.ModuleDuration,
		//Runtime:        input.Runtime,
		ExamType: input.ExamType,
	}

	moduleValidator := validator.New()

	if data.ValidateModuleInfo(moduleValidator, moduleInfo); !moduleValidator.Valid() {
		app.failedValidationResponse(w, r, moduleValidator.Errors)
		return
	}

	err = app.models.ModuleInfo.Insert(moduleInfo)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/noduleinfo/%d", moduleInfo.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"module_info": moduleInfo}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showModuleInfoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDPAram(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	moduleInfo, err := app.models.ModuleInfo.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"module_info": moduleInfo}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) UpdateModuleInfoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDPAram(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	moduleInfo, err := app.models.ModuleInfo.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		ModuleName     string   `json:"module_name"`
		ModuleDuration int32    `json:"module_duration"`
		ExamType       []string `json:"exam_type"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	moduleInfo.ModuleName = input.ModuleName
	moduleInfo.ModuleDuration = input.ModuleDuration
	moduleInfo.ExamType = input.ExamType

	v := validator.New()

	if data.ValidateModuleInfo(v, moduleInfo); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.ModuleInfo.Update(moduleInfo)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"module_info": moduleInfo}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) DeleteModuleInfoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDPAram(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.ModuleInfo.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "ModuleInfo successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
