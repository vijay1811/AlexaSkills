package alexa

import "AlexaSkills/protocol/alexa"

type RequestHandler interface {
	HandleRequest(r *alexa.AlexaRequest) (*alexa.AlexaResponse, error)
}
