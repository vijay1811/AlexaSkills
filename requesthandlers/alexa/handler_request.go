package alexa

import (
	"errors"

	"AlexaSkills/protocol/alexa"
	"AlexaSkills/requesthandlers/alexa/application/travelAssistant"
)

const (
	aapId_travelAssistant = "amzn1.ask.skill.4a81bcdd-bae1-4330-9f04-10ef53e58e10"
)

type Default struct {
	aaps map[string]RequestHandler
}

func (d *Default) HandleRequest(r *alexa.AlexaRequest) (*alexa.AlexaResponse, error) {
	h, ok := d.aaps[r.Session.Application.ApplicationID]
	if !ok {
		return nil, errors.New("Application is not supported")
	}
	return h.HandleRequest(r)
}

func NewAlexaRequestHandler() RequestHandler {
	aaps := make(map[string]RequestHandler)
	aaps[aapId_travelAssistant] = travelAssistant.NewRequestHandler()
	return &Default{
		aaps: aaps,
	}
}
