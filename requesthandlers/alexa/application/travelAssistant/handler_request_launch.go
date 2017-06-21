package travelAssistant

import "AlexaSkills/protocol/alexa"

type handlerLaunchRequest struct {
}

func (handlerLaunchRequest) handleRequest(r *alexa.AlexaRequest) (*alexa.AlexaResponse, error) {
	return &alexa.AlexaResponse{
		Version: r.Version,
		Response: &alexa.Response{
			OutputSpeech: &alexa.OutputSpeech{
				Type: "PlainText",
				Text: "Hi, This is your Travel Assistant. I can help you book a cab.",
			},
		},
	}, nil
}
