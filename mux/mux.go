package mux

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	muxAlexa "AlexaSkills/mux/alexa"
	alexaProtocol "AlexaSkills/protocol/alexa"
	rhAlexa "AlexaSkills/requesthandlers/alexa"
)

type Handler struct {
	handlers map[string]RequestHandler
	rh       rhAlexa.RequestHandler
	slots    []*alexaProtocol.Slot
}

func NewHandler() *Handler {
	h := &Handler{
		handlers: make(map[string]RequestHandler),
		rh:       rhAlexa.NewAlexaRequestHandler(),
	}

	// Multiple handlers
	h.handlers["/alexa"] = &muxAlexa.Handler{}

	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("HTTP Request Received: \n%+v\nMethod: %+v\nURL: %+v\nBODY: %+v", r, r.Method, r.URL, r.Body)

	log.Printf("Url: %s\n", r.URL.String())
	for key, handler := range h.handlers {
		if strings.HasPrefix(r.URL.String(), key) {
			log.Printf("got handler key: %s\n", key)
			handler.ServeHTTP(w, r, h.rh)
			return
		}
	}
	log.Printf("Url: %s\n", r.URL.String())
	io.WriteString(w, "My server: "+r.URL.String())
}
