package travelAssistant

import (
	"errors"
	"log"

	"AlexaSkills/protocol/alexa"
)

type Default struct {
	handlers map[alexa.RequestType]requestHandler
}

func NewRequestHandler() *Default {
	handlers := make(map[alexa.RequestType]requestHandler)
	handlers[alexa.RequestType_IntentRequest] = handlerIntentRequest{}
	handlers[alexa.RequestType_LaunchRequest] = handlerLaunchRequest{}
	handlers[alexa.RequestType_SessionEndedRequestt] = handlerSessionEndedRequest{}
	return &Default{
		handlers: handlers,
	}
}

func (d *Default) HandleRequest(r *alexa.AlexaRequest) (*alexa.AlexaResponse, error) {
	h, ok := d.handlers[r.Request.Type]
	if !ok {
		log.Printf("request Type : %v , type: %T", r.Request.Type, r.Request.Type)
		return nil, errors.New("request not supported")
	}
	return h.handleRequest(r)
}
