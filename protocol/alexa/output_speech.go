package alexa

type OutputSpeech struct {
	Type string `json:"type"`
	Text string `json:"text"`
	SSML string `json:"ssml"`
}
