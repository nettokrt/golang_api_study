package main

import (
	"net/http"
)

func (app *application) healthCheckHanler( w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status": "ok",
		"env": app.config.env,
		"version": version,
		"message": "API is running",
  }
	
	if err := writejSON(w, http.StatusOK, data); err != nil {
		writeJSONError(w, http.StatusInternalServerError, "err.Error()")
	}
}

