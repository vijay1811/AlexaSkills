package alexa

import (
	"errors"

	"github.com/personalbuildhelper/protocol/alexa"
	"github.com/personalbuildhelper/requesthandlers/alexa/application/buildhelper"
)

const (
	aapId_buildHelper = "amzn1.ask.skill.462a6bc5-8525-48cd-9b6a-bc862fc1b667"
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
	aaps[aapId_buildHelper] = buildhelper.NewRequestHandler()
	return &Default{
		aaps: aaps,
	}
}
