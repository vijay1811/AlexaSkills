package main

import (
	"net/http"
	"os"

	"AlexaSkills/mux"
)

/*var (
	paramChan chan *request
	ffmChan   chan []byte
)

func init() {
	paramChan = make(chan *request, 20)
	ffmChan = make(chan []byte, 20)
}
*/

/*func hello(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("ERROR: %v\n", err)
		return
	}
	log.Printf("method: %v\n", r.Method)
	log.Printf("URL: %+v\n", r.URL)
	log.Printf("request body: %s\n", body)

	header := w.Header()
	header.Add("Content-type", "application/json")
	req := &request{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Printf("ERROR: %v\n", err)
		return
	}
	req.Result.ServiceType = "Google Home"
	paramChan <- req
	ffm := <-ffmChan
	log.Printf("received fullfillment: %s\n", ffm)
	w.Write(ffm)
}
*/

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	server := http.Server{
		Addr:    ":" + port,
		Handler: mux.NewHandler(),
	}

	server.ListenAndServe()
}

/*func handleSystemClient(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-type", "application/json")
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
	w.Write(data)
}
*/
