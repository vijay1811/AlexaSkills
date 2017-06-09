package mux

import (
	"net/http"

	"AlexaSkills/requesthandlers/alexa"
)

type RequestHandler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request, rh alexa.RequestHandler)
}
