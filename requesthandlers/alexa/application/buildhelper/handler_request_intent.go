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
		if _, exists := slots[key]; !exists {
			slots[key] = attributes[key]
		}
	}

	buildType, buildTypeGiven := slots["buildtypeslot"]
	sourceType, sourceTypeGiven := slots["buildsourcetypeslot"]

	switch {
	case buildTypeGiven && sourceTypeGiven:
		outSpeech = &alexa.OutputSpeech{
			Type: "PlainText",
			Text: fmt.Sprintf("%s is build for %s", sourceType.Value, buildType.Value),
		}
		isComplete = true
	case !buildTypeGiven && sourceTypeGiven:
		outSpeech = &alexa.OutputSpeech{
			Type: "PlainText",
			Text: "Please tell the build type",
		}
	case buildTypeGiven && !sourceTypeGiven:
		outSpeech = &alexa.OutputSpeech{
			Type: "PlainText",
			Text: "Please tell the source type",
		}
	case !buildTypeGiven && !sourceTypeGiven:
		outSpeech = &alexa.OutputSpeech{
			Type: "PlainText",
			Text: "Please tell source and type for build",
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
