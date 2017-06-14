package buildhelper

import (
	"fmt"
	"log"

	"AlexaSkills/protocol/alexa"
)

type handlerIntentRequest struct {
}

func (handlerIntentRequest) handleRequest(r *alexa.AlexaRequest) (*alexa.AlexaResponse, error) {

	var resp *alexa.AlexaResponse
	attributes := make(map[string]*alexa.Slot)

	if r.Session.Attributes != nil {
		attributes = r.Session.Attributes
	}

	log.Printf("ATTRIBUTES RECEIVED IN HANDLER INTENT:\n%+v", attributes)

	outputSpeech, slots, isComplete := getOutputSpeech(r.Request.Intent, attributes)

	if isComplete {
		resp = &alexa.AlexaResponse{
			Version: r.Version,
			Response: &alexa.Response{
				OutputSpeech:     outputSpeech,
				ShouldEndSession: true,
			},
		}
	} else {
		resp = &alexa.AlexaResponse{
			Version:           r.Version,
			SessionAttributes: slots,
			Response: &alexa.Response{
				OutputSpeech:     outputSpeech,
				ShouldEndSession: false,
			},
		}
	}

	return resp, nil
}

func getOutputSpeech(intent *alexa.Intent, attributes map[string]*alexa.Slot) (*alexa.OutputSpeech, map[string]*alexa.Slot, bool) {

	var outSpeech *alexa.OutputSpeech
	var isComplete bool

	// All the newly requested slots in the intent must be saved in a new map.
	slots := intent.Slots

	for slotName := range attributes {
		intentSlot := slots[slotName]
		log.Printf("Intent Slot: %+v", intentSlot)

		attributeSlot := attributes[slotName]
		log.Printf("Attribute Slot: %+v", attributeSlot)

		if intentSlot.Value == "" && attributeSlot.Value != "" {
			delete(slots, slotName)
			slots[slotName] = attributeSlot
		}
	}

	buildTypeSlot := slots["buildtypeslot"]
	buildSourceTypeSlot := slots["buildsourcetypeslot"]

	var buildType, sourceType string
	var buildTypeGiven, sourceTypeGiven bool

	if buildTypeSlot.Value != "" {
		buildType = buildTypeSlot.Value
		buildTypeGiven = true
	}

	if buildSourceTypeSlot.Value != "" {
		sourceType = buildSourceTypeSlot.Value
		sourceTypeGiven = true
	}

	switch {
	case buildTypeGiven && sourceTypeGiven:
		outSpeech = &alexa.OutputSpeech{
			Type: "PlainText",
			Text: fmt.Sprintf("'%s' is build for '%s'", sourceType, buildType),
		}
		isComplete = true
	case !buildTypeGiven && sourceTypeGiven:
		outSpeech = &alexa.OutputSpeech{
			Type: "PlainText",
			Text: fmt.Sprintf("Please tell the build type for build source '%s'", sourceType),
		}
	case buildTypeGiven && !sourceTypeGiven:
		outSpeech = &alexa.OutputSpeech{
			Type: "PlainText",
			Text: fmt.Sprintf("Please tell the build source type for the build type '%s'", buildType),
		}
	case !buildTypeGiven && !sourceTypeGiven:
		outSpeech = &alexa.OutputSpeech{
			Type: "PlainText",
			Text: "Please tell the build source type and build type",
		}
	default:
		outSpeech = &alexa.OutputSpeech{
			Type: "PlainText",
			Text: "This is not possible",
		}
	}

	return outSpeech, slots, isComplete
}
