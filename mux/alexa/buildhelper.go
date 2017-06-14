package alexa

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	alexaProtocol "AlexaSkills/protocol/alexa"
	"AlexaSkills/requesthandlers/alexa"
)

func init() {
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

	log.Printf("ALEXA REQUEST RECEIVED: %+v", alexaReq)
	log.Printf("ALEXA REQUEST SESSION OBJECT RECEIVED: %+v", alexaReq.Session)
	log.Printf("ALEXA REQUEST SESSION ATTRIBUTES RECEIVED: %+v", alexaReq.Session.Attributes)

	alexaResp, err := rh.HandleRequest(alexaReq)
	if err != nil {
		// TODO handle errors gracefully here
		log.Printf("ERROR: %v\n", err)
		return
	}

	log.Printf("SENDING ALEXA RESPONSE: %+v", alexaResp)

	resp, err := json.Marshal(alexaResp)
	if err != nil {
		log.Printf("ERROR: %v\n", err)
		return
	}

	log.Printf("RESPONSE: %v\n", string(resp))
	w.Write(resp)
}
