package alexa

type Intent struct {
	Name  string           `json,"name"`
	Slots map[string]*Slot `json,"slots"`
}

type Slot struct {
	Name  string
	Value string
}
