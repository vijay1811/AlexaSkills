package mux

import (
	"io"
	"log"
	"net/http"
	"strings"

	muxAlexa "github.com/personalbuildhelper/mux/alexa"
	"github.com/personalbuildhelper/mux/googlehome"
	"github.com/personalbuildhelper/mux/system"
	rhAlexa "github.com/personalbuildhelper/requesthandlers/alexa"
)

type Handler struct {
	handlers map[string]RequestHandler
	rh       rhAlexa.RequestHandler
}

func NewHandler() *Handler {
	h := &Handler{
		handlers: make(map[string]RequestHandler),
		rh:       rhAlexa.NewAlexaRequestHandler(),
	}
	h.handlers["/alexa"] = &muxAlexa.Handler{}
	h.handlers["/gh"] = &googlehome.Handler{}
	h.handlers["/system"] = &system.Handler{}
	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

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
