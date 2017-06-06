package mux

import (
	"net/http"

	"github.com/personalbuildhelper/requesthandlers/alexa"
)

type RequestHandler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request, rh alexa.RequestHandler)
}
