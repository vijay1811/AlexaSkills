package alexa

import "github.com/personalbuildhelper/protocol/alexa"

type RequestHandler interface {
	HandleRequest(r *alexa.AlexaRequest) (*alexa.AlexaResponse, error)
}
