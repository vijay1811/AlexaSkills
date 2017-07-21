package alexa

import (
	"AlexaSkills/protocol/alexa"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type RequestHandler interface {
	HandleRequest(r *alexa.AlexaRequest, cl mqtt.Client) (*alexa.AlexaResponse, error)
}
