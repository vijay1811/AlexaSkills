package alexa

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	alexaProtocol "github.com/personalbuildhelper/protocol/alexa"
	"github.com/personalbuildhelper/requesthandlers/alexa"
)

func init() {
	mux = make(map[string]func(http.ResponseWriter, *http.Request, alexa.RequestHandler))
	mux["/alexa/buildhelper"] = alexaBuildHelper
}

func alexaBuildHelper(w http.ResponseWriter, r *http.Request, rh alexa.RequestHandler) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("ERROR: %v\n", err)
		return
	}
	log.Printf("method: %v\n", r.Method)
	log.Printf("URL: %+v\n", r.URL)
	log.Printf("request body: %s\n", body)

	alexaReq := &alexaProtocol.AlexaRequest{}
	err = json.Unmarshal(body, &alexaReq)
	if err != nil {
		log.Printf("ERROR: %v\n", err)
		return
	}

	alexaResp, err := rh.HandleRequest(alexaReq)
	if err != nil {
		// TODO handle errors gracefully here
		log.Printf("ERROR: %v\n", err)
		return
	}

	resp, err := json.Marshal(alexaResp)
	if err != nil {
		log.Printf("ERROR: %v\n", err)
		return
	}
	w.Write(resp)
}
