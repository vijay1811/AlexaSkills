package travelAssistant

import "AlexaSkills/protocol/alexa"

type requestHandler interface {
	handleRequest(r *alexa.AlexaRequest) (*alexa.AlexaResponse, error)
}
