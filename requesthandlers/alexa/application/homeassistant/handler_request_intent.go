package buildhelper

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	"AlexaSkills/protocol/alexa"
)

type handlerIntentRequest struct {
}

func (handlerIntentRequest) handleRequest(r *alexa.AlexaRequest, cl mqtt.Client) (*alexa.AlexaResponse, error) {

	var resp *alexa.AlexaResponse
	attributes := make(map[string]*alexa.Slot)

	if r.Session.Attributes != nil {
		attributes = r.Session.Attributes
	}

	log.Printf("ATTRIBUTES RECEIVED IN HANDLER INTENT:\n%+v", attributes)

	outputSpeech, slots, isComplete := getOutputSpeech(r.Request.Intent, attributes, cl)

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

func getOutputSpeech(intent *alexa.Intent, attributes map[string]*alexa.Slot, cl mqtt.Client) (*alexa.OutputSpeech, map[string]*alexa.Slot, bool) {

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

		action = strings.Title(strings.ToLower(action))
		device = strings.Title(strings.ToLower(device))
		location = strings.Title(strings.ToLower(location))

		go publishToMQTT(cl, action, location, device)

		outSpeech = &alexa.OutputSpeech{
			Type: "PlainText",
			// Text: fmt.Sprintf("I did %s the %s's %s for you.", action, location, device),
			Text: "In Progress",
		}
		isComplete = true

		// if err := publishToMQTT(cl, action, location, device); err != nil {
		// 	outSpeech = &alexa.OutputSpeech{
		// 		Type: "PlainText",
		// 		Text: "There seems to be some problem communicating over MQTT",
		// 	}
		// } else {
		// 	outSpeech = &alexa.OutputSpeech{
		// 		Type: "PlainText",
		// 		// Text: fmt.Sprintf("I did %s the %s's %s for you.", action, location, device),
		// 		Text: "In Progress",
		// 	}
		// 	isComplete = true
		// }
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

func publishToMQTT(mqttCl mqtt.Client, action, location, device string) {
	fmt.Printf("Connecting the 'Client_Alexa' to publish\n")
	if token := mqttCl.Connect(); token.Wait() && token.Error() != nil {
		err := token.Error()
		fmt.Printf("Error Connecting MQTT Client: %v\n", err.Error())
		return
	}

	if strings.Contains(strings.ToLower(action), "on") || strings.ToLower(action) == "start" {
		action = "On"
	} else {
		action = "Off"
	}

	skillMap := map[string]string{
		"Action":   action,
		"Location": location,
		"Device":   device,
	}

	payload, err := json.Marshal(skillMap)
	if err != nil {
		fmt.Printf("Error Marshalling the skillMap: %v", err.Error())
		return
	}

	// Publishing New Information
	fmt.Printf("Publishing Message over '/homeAutomation/%s' topic\n", device)
	token := mqttCl.Publish(fmt.Sprintf("/homeAutomation/%s", device), 0, false, payload)
	token.Wait()

	fmt.Printf("Disconnecting 'Client_Alexa'\n")
	mqttCl.Disconnect(250)

	return
}
