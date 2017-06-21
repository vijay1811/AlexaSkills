package alexa

type AlexaResponse struct {
	Version           string           `json:"version"`
	SessionAttributes map[string]*Slot `json:"sessionAttributes,omitempty"`
	Response          *Response        `json:"response"`
}

type Response struct {
	OutputSpeech *OutputSpeech `json:"outputSpeech"`
	// Card             *Card         `json:"card"`
	ShouldEndSession bool `json:"shouldEndSession"`
}

// type Card struct {
// 	Type *CardType `json:"type"`
// }

// type CardType string

// const (
// 	CardType_Simple      CardType = "Simple"
// 	CardType_Standard    CardType = "Standard"
// 	CardType_LinkAccount CardType = "LinkAccount"
// )
