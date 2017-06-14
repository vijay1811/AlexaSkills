package buildhelper

import (
	"fmt"

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

	for key := range attributes {
		if slots[key].Value != "" && attributes[key].Value != "" {
			slots[key].Value = attributes[key].Value
		}
	}

	// for key := range attributes {
	// 	if _, exists := slots[key]; !exists {
	// 		slots[key] = attributes[key]
	// 	}
	// }

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

	// buildType, buildTypeGiven := slots["buildtypeslot"]
	// sourceType, sourceTypeGiven := slots["buildsourcetypeslot"]

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

/*amzn1.ask.skill.462a6bc5-8525-48cd-9b6a-bc862fc1b667
amzn1.ask.skill.462a6bc5-8525-48cd-9b6a-bc862fc1b667*/
