package buildhelper

import "AlexaSkills/protocol/alexa"

type handlerLaunchRequest struct {
}

func (handlerLaunchRequest) handleRequest(r *alexa.AlexaRequest) (*alexa.AlexaResponse, error) {
	return &alexa.AlexaResponse{
		Version: r.Version,
		Response: &alexa.Response{
			OutputSpeech: &alexa.OutputSpeech{
				Type: "PlainText",
				Text: "Hi, I am your Home Assistant. How can I help you?",
			},
		},
	}, nil
}
