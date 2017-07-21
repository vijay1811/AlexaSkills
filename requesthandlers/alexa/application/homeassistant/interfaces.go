package buildhelper

import (
	"AlexaSkills/protocol/alexa"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type requestHandler interface {
	handleRequest(r *alexa.AlexaRequest, cl mqtt.Client) (*alexa.AlexaResponse, error)
}
