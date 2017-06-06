package buildhelper

import "github.com/personalbuildhelper/protocol/alexa"

type handlerLaunchRequest struct {
}

func (handlerLaunchRequest) handleRequest(r *alexa.AlexaRequest) (*alexa.AlexaResponse, error) {
	return &alexa.AlexaResponse{
		Version: r.Version,
		Response: &alexa.Response{
			OutputSpeech: &alexa.OutputSpeech{
				Type: "PlainText",
				Text: "Hi I am build helper I will help you to build",
			},
		},
	}, nil
}
