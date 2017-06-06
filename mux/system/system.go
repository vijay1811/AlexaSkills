package system

import (
	"net/http"

	"github.com/personalbuildhelper/requesthandlers/alexa"
)

func handleSystemClient(w http.ResponseWriter, r *http.Request, rh alexa.RequestHandler) {

	/*w.Header().Add("Content-type", "application/json")
	timer := time.NewTimer(20 * time.Second)
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("System Client ERROR: %v\n", err)
			return
		}

		log.Printf("Before sending body\n")
		ffmChan <- body
		log.Printf("After sending body\n")
		return
	}
	var req *request
	select {
	case req = <-paramChan:
	case <-timer.C:
		req = &request{
			Result: &Result{
				Parameters: &Parameters{
					Status: -1,
				},
			},
		}

	}
	data, err := json.Marshal(req)
	if err != nil {
		log.Printf("ERROR: %v\n", err)
		return
	}
	log.Printf("response body: %s\n", data)
	w.Write(data)*/
}
