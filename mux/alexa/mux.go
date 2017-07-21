package alexa

import (
	"io"
	"net/http"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	"AlexaSkills/requesthandlers/alexa"
)

var mux = make(map[string]func(http.ResponseWriter, *http.Request, alexa.RequestHandler, mqtt.Client))

type Handler struct {
}

func (Handler) ServeHTTP(w http.ResponseWriter, r *http.Request, rh alexa.RequestHandler, cl mqtt.Client) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r, rh, cl)
		return
	}
	io.WriteString(w, "My server: "+r.URL.String())
}
