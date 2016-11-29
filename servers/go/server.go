package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/gorilla/mux"
)

//DummyInfo the info that will be sent by the client
type DummyInfo struct {
	DummyString string
	DummyFloat  float64
	DummyInt    int
	DummyDate   time.Time
}

//ProcessorHandler process received message
func ProcessorHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var dmi DummyInfo
	err = json.Unmarshal(data, &dmi)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dmi)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/bombhere", ProcessorHandler).Methods("POST")

	http.ListenAndServe(":10000", r)
	time.Sleep(65 * time.Second)
	log.Println("Exiting")
}
