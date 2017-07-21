package mux

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	muxAlexa "AlexaSkills/mux/alexa"
	rhAlexa "AlexaSkills/requesthandlers/alexa"
)

type Handler struct {
	handlers   map[string]RequestHandler
	rh         rhAlexa.RequestHandler
	mqttClient mqtt.Client
	// slots    []*alexaProtocol.Slot
}

func NewHandler() *Handler {

	h := &Handler{
		handlers:   make(map[string]RequestHandler),
		rh:         rhAlexa.NewAlexaRequestHandler(),
		mqttClient: createNewClient(),
	}

	// Multiple handlers
	h.handlers["/alexa"] = &muxAlexa.Handler{}

	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// fmt.Printf("\nHTTP Request Received: \n%+v\nMETHOD: %+v\nURL: %+v", r, r.Method, r.URL)

	log.Printf("Url: %s\n", r.URL.String())
	for key, handler := range h.handlers {
		if strings.HasPrefix(r.URL.String(), key) {
			log.Printf("got handler key: %s\n", key)
			handler.ServeHTTP(w, r, h.rh, h.mqttClient)
			return
		}
	}
	log.Printf("Url: %s\n", r.URL.String())
	io.WriteString(w, "My server: "+r.URL.String())
}

func createNewClient() mqtt.Client {
	fmt.Printf("Setting up Client for Alexa Skill (Client_Alexa)\n")
	fmt.Printf("Setting up Options for 'Client_Alexa'\n")
	options := mqtt.NewClientOptions().AddBroker("tcp://iot.eclipse.org:1883").SetClientID("Client_Alexa")

	fmt.Printf("Setting up 'Client_Alexa'\n")
	alexa_client := mqtt.NewClient(options)

	return alexa_client

	// fmt.Printf("Connecting the 'Client_Alexa'\n")
	// if token := alexa_client.Connect(); token.Wait() && token.Error() != nil {
	// 	panic(token.Error())
	// }

	// fmt.Printf("Disconnecting 'Client_Alexa'\n")
	// alexa_client.Disconnect(250)
}
