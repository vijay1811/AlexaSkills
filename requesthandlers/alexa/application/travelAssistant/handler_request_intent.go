package travelAssistant

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

	sourceSlot := slots["SourceSlot"]
	destinationSlot := slots["DestinationSlot"]

	var source, destination string
	var sourceGiven, destinationGiven bool

	if sourceSlot.Value != "" {
		source = sourceSlot.Value
		sourceGiven = true
	}

	if destinationSlot.Value != "" {
		destination = destinationSlot.Value
		destinationGiven = true
	}

	switch {
	case sourceGiven && destinationGiven:
		outSpeech = &alexa.OutputSpeech{
			Type: "PlainText",
			Text: fmt.Sprintf("Your taxi is being booked for travel from '%s' to '%s'. Enjoy Your Ride!", source, destination),
		}
		isComplete = true
	case !sourceGiven && destinationGiven:
		outSpeech = &alexa.OutputSpeech{
			Type: "PlainText",
			Text: fmt.Sprintf("Please specify the starting point for the journey to '%s'", destination),
		}
	case sourceGiven && !destinationGiven:
		outSpeech = &alexa.OutputSpeech{
			Type: "PlainText",
			Text: fmt.Sprintf("Please specify the destination for your journey from '%s'", source),
		}
	case !sourceGiven && !destinationGiven:
		outSpeech = &alexa.OutputSpeech{
			Type: "PlainText",
			Text: "Can I know the source or the destination for the ride?",
		}
	default:
		outSpeech = &alexa.OutputSpeech{
			Type: "PlainText",
			Text: "This is not possible",
		}
	}

	return outSpeech, slots, isComplete
}
