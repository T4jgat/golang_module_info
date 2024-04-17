package main

import "net/http"

func (app *application) ShowAllTeachersHandler(w http.ResponseWriter, r *http.Request) {
	teachers, err := app.models.TeacherInfo.GetAll()
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"teachers": teachers}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
