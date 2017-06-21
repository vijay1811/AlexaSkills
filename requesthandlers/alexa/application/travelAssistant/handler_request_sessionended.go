package travelAssistant

import (
	"errors"

	"AlexaSkills/protocol/alexa"
)

type handlerSessionEndedRequest struct {
}

func (handlerSessionEndedRequest) handleRequest(r *alexa.AlexaRequest) (*alexa.AlexaResponse, error) {
	return nil, errors.New("Session Ended")
}
