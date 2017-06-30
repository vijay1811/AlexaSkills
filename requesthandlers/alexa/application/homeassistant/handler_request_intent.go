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
	locationSlot := slots["locationSlot"]

	var action, device, location string
	var actionGiven, deviceGiven, locationGiven bool

	if actionSlot.Value != "" {
		action = actionSlot.Value
		actionGiven = true
	}

	if deviceSlot.Value != "" {
		device = deviceSlot.Value
		deviceGiven = true
	}

	if locationSlot.Value != "" {
		location = locationSlot.Value
		locationGiven = true
	}

	switch {
	case actionGiven && deviceGiven && locationGiven:
		outSpeech = &alexa.OutputSpeech{
			Type: "PlainText",
			Text: fmt.Sprintf("I did %s the %s's %s for you.", action, location, device),
		}
		isComplete = true
	case actionGiven && deviceGiven && !locationGiven:
		outSpeech = &alexa.OutputSpeech{
			Type: "PlainText",
			Text: fmt.Sprintf("Please specify the location of your installed %s.", device),
		}
	case actionGiven && !deviceGiven && locationGiven:
		outSpeech = &alexa.OutputSpeech{
			Type: "PlainText",
			Text: fmt.Sprintf("Ok! I am in %s. Please specify the device to %s", location, action),
		}
	case actionGiven && !deviceGiven && !locationGiven:
		outSpeech = &alexa.OutputSpeech{
			Type: "PlainText",
			Text: fmt.Sprintf("Please specify the location and the device to %s", action),
		}
	case !actionGiven && deviceGiven && locationGiven:
		outSpeech = &alexa.OutputSpeech{
			Type: "PlainText",
			Text: fmt.Sprintf("What would you want me to do with %s's %s", location, device),
		}
	case !actionGiven && deviceGiven && !locationGiven:
		outSpeech = &alexa.OutputSpeech{
			Type: "PlainText",
			Text: fmt.Sprintf("Please specify the location of your installed %s", device),
		}
	case !actionGiven && !deviceGiven && locationGiven:
		outSpeech = &alexa.OutputSpeech{
			Type: "PlainText",
			Text: fmt.Sprintf("Ok! I am in %s. Please specify the device", location),
		}
	case !actionGiven && !deviceGiven && !locationGiven:
		outSpeech = &alexa.OutputSpeech{
			Type: "PlainText",
			Text: fmt.Sprintf("I could be helpful if I know what to do. Let's begin with device or location"),
		}
	default:
		outSpeech = &alexa.OutputSpeech{
			Type: "PlainText",
			Text: "This is not possible",
		}
	}

	return outSpeech, slots, isComplete
}
