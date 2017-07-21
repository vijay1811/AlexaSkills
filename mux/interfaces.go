package mux

import (
	"net/http"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	"AlexaSkills/requesthandlers/alexa"
)

type RequestHandler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request, rh alexa.RequestHandler, cl mqtt.Client)
}
