package alexa

import (
	"errors"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	"AlexaSkills/protocol/alexa"
	assistant "AlexaSkills/requesthandlers/alexa/application/homeassistant"
)

const (
	aapId_buildHelper = "amzn1.ask.skill.1df8d7cb-77c9-4cb7-839e-eb1847849e8d"
)

type Default struct {
	aaps map[string]RequestHandler
}

func (d *Default) HandleRequest(r *alexa.AlexaRequest, cl mqtt.Client) (*alexa.AlexaResponse, error) {
	h, ok := d.aaps[r.Session.Application.ApplicationID]
	if !ok {
		return nil, errors.New("Application is not supported")
	}
	return h.HandleRequest(r, cl)
}

func NewAlexaRequestHandler() RequestHandler {
	aaps := make(map[string]RequestHandler)
	aaps[aapId_buildHelper] = assistant.NewRequestHandler()
	return &Default{
		aaps: aaps,
	}
}
