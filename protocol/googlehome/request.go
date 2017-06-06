package googlehome

type request struct {
	Result *Result `json:"result"`
}

type Result struct {
	ResolvedQuery string      `json:"resolvedQuery"`
	MetaData      *MetaData   `json:"metadata"`
	Parameters    *Parameters `json:"parameters"`
}

type MetaData struct {
	IntentName string `json:"intentName"`
	IntenID    string
}

type Parameters struct {
	Status   int
	Action   string
	Location string
}
