package alexa

type AlexaResponse struct {
	Version  string    `json:"version"`
	Response *Response `json:"response"`
}

type Response struct {
	OutputSpeech *OutputSpeech `json:"outputSpeech"`
}
