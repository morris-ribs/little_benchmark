package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

//DummyInfo the info that will be sent by the client
type DummyInfo struct {
	DummyString string
	DummyFloat  float64
	DummyInt    int
	DummyDate   time.Time
}

func main() {
	log.SetFlags(0)

	duration := flag.Int("duration", 60, "Duration to run in seconds.")
	flag.Parse()

	var data []byte
	var err error

	if err != nil {
		log.Fatalf("Failed to create request body: %v", err)
	}

	log.Printf("Sending requests for %d seconds. Request size: %d\n", *duration, len(data))

	requestTotal := 0
	go func() {
		for {
			var dmi DummyInfo
			dmi.DummyString = "Dummy string " + strconv.Itoa(requestTotal)
			dmi.DummyInt = requestTotal
			dmi.DummyFloat = rand.Float64()
			dmi.DummyDate = time.Now()

			dataToSend, err := json.Marshal(&dmi)

			// send the data to the server
			req, err := http.NewRequest("POST", "http://127.0.0.1:11000/bombhere", bytes.NewBuffer(dataToSend))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			if err != nil {
				log.Fatalf("Failed http request: %v", err)
			}
			if resp.StatusCode != http.StatusOK {
				log.Fatalf("Error response code: %v", resp.StatusCode)
			}

			requestTotal++
		}
	}()

	time.Sleep(time.Duration(*duration) * time.Second)
	log.Printf("Request Total: %d\n", requestTotal)
}
