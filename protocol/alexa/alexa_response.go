package alexa

type AlexaResponse struct {
	Version           string           `json:"version"`
	SessionAttributes map[string]*Slot `json:"sessionAttributes,omitempty"`
	Response          *Response        `json:"response"`
}

type Response struct {
	OutputSpeech     *OutputSpeech `json:"outputSpeech"`
	ShouldEndSession bool          `json:"shouldEndSession"`
}
