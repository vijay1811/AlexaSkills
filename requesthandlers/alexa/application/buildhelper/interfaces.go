package buildhelper

import "github.com/personalbuildhelper/protocol/alexa"

type requestHandler interface {
	handleRequest(r *alexa.AlexaRequest) (*alexa.AlexaResponse, error)
}
