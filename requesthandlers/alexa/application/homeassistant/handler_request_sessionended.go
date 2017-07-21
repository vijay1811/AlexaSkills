package buildhelper

import (
	"errors"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	"AlexaSkills/protocol/alexa"
)

type handlerSessionEndedRequest struct {
}

func (handlerSessionEndedRequest) handleRequest(r *alexa.AlexaRequest, cl mqtt.Client) (*alexa.AlexaResponse, error) {
	return nil, errors.New("Session Ended")
}
