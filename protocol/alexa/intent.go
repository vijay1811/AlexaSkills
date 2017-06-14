package alexa

type Intent struct {
	Name  string           `json,"name"`
	Slots map[string]*Slot `json,"slots"`
}

type Slot struct {
	Name  string `json,"name"`
	Value string `json,"value"`
	// ConfirmationStatus ConfirmationStatus `json,"confirmationstatus"`
}

// type ConfirmationStatus string

// const (
// 	ConfirmationStatus_NONE      ConfirmationStatus = "NONE"
// 	ConfirmationStatus_CONFIRMED ConfirmationStatus = "CONFIRMED"
// 	ConfirmationStatus_DENIED    ConfirmationStatus = "DENIED"
// )

type IntentName string

const (
	IntentName_Build IntentName = "build"
	IntentName_Test  IntentName = "test"
)

// const (
// 	slots_BuilderHelper = []string{"buildtypeslot", "buildsourcetypeslot"}
// )
