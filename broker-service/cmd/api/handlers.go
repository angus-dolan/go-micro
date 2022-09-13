package main

import (
	"net/http"
)

func (app *Config) Broker(w http.ResponseWriter, req *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Broker used",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}
