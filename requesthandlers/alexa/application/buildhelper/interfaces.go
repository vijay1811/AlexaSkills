package buildhelper

import "AlexaSkills/protocol/alexa"

type requestHandler interface {
	handleRequest(r *alexa.AlexaRequest) (*alexa.AlexaResponse, error)
}
