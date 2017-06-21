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

	actionSlot := slots["actionSlot"]
	deviceSlot := slots["deviceSlot"]

	var action, device string
	var actionGiven, deviceGiven bool

	if actionSlot.Value != "" {
		action = actionSlot.Value
		actionGiven = true
	}

	if deviceSlot.Value != "" {
		device = deviceSlot.Value
		deviceGiven = true
	}

	switch {
	case actionGiven && deviceGiven:
		outSpeech = &alexa.OutputSpeech{
			Type: "PlainText",
			Text: fmt.Sprintf("I did %s the %s for you.", action, device),
		}
		isComplete = true
	case !actionGiven && deviceGiven:
		outSpeech = &alexa.OutputSpeech{
			Type: "PlainText",
			Text: fmt.Sprintf("Please specify the action on the device %s", device),
		}
	case actionGiven && !deviceGiven:
		outSpeech = &alexa.OutputSpeech{
			Type: "PlainText",
			Text: fmt.Sprintf("Please specify which device to perform action %s on", action),
		}
	case !actionGiven && !deviceGiven:
		outSpeech = &alexa.OutputSpeech{
			Type: "PlainText",
			Text: "I could be helpful if I know what to do. Please specify action or device.",
		}
	default:
		outSpeech = &alexa.OutputSpeech{
			Type: "PlainText",
			Text: "This is not possible",
		}
	}

	return outSpeech, slots, isComplete
}
