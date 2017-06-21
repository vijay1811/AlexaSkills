package alexa

type RequestType string

const (
	RequestType_LaunchRequest        RequestType = "LaunchRequest"
	RequestType_IntentRequest        RequestType = "IntentRequest"
	RequestType_SessionEndedRequestt RequestType = "SessionEndedRequest"
)

type AlexaRequest struct {
	Version string   `json:"version"`
	Session *Session `json:"session"`
	Request *Request `json:"request"`
}

type Session struct {
	Application *Application     `json:"application"`
	Attributes  map[string]*Slot `json:"attributes,omitempty"`
	User        *User            `json:"user"`
}

type Application struct {
	ApplicationID string `json:"applicationId"`
}

type User struct {
	UserID      string `json:"userId"`
	AccessToken string `json:"accessToken"`
}

type Request struct {
	Type   RequestType `json:"type"`
	Reason string      `json:"reason"`
	Intent *Intent     `json:"intent:`
}
