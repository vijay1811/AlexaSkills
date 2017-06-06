package alexa

import (
	"io"
	"net/http"

	"github.com/personalbuildhelper/requesthandlers/alexa"
)

var mux = make(map[string]func(http.ResponseWriter, *http.Request, alexa.RequestHandler))

type Handler struct {
}

func (Handler) ServeHTTP(w http.ResponseWriter, r *http.Request, rh alexa.RequestHandler) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r, rh)
		return
	}
	io.WriteString(w, "My server: "+r.URL.String())
}
