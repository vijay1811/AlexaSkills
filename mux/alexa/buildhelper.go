package alexa

import (
	"encoding/json"
	"fmt"
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

	fmt.Printf("HTTP Request Received: %+v", r)
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

	fmt.Printf("Alexa Request Received: %+v", alexaReq)

	alexaResp, err := rh.HandleRequest(alexaReq)
	if err != nil {
		// TODO handle errors gracefully here
		log.Printf("ERROR: %v\n", err)
		return
	}

	fmt.Printf("Sending Alexa Response: %+v", alexaResp)

	resp, err := json.Marshal(alexaResp)
	if err != nil {
		log.Printf("ERROR: %v\n", err)
		return
	}
	w.Write(resp)
}
