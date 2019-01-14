package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Status int

const (
	FAILING = 0
	RUNNING = 1
)

type MockResponse struct {
	SystemStatus Status `json:"systemStatus"`
	Env          string `json:"env"`
	ID           int    `json:"id"`
}

func main() {
	http.HandleFunc("/metrics", getMetrics)
	log.Println("Serving on 8010..")
	log.Fatal(http.ListenAndServe(":8010", nil))
}

func getMetrics(w http.ResponseWriter, r *http.Request) {
	metrics := []MockResponse{}
	var env string
	var status Status
	for i := 0; i < 6; i++ {
		if i%2 == 0 {
			env = "prod"
		} else {
			env = "dev"
		}
		if time.Now().UnixNano()%2 == 0 {
			status = FAILING
		} else {
			status = RUNNING
		}
		metrics = append(metrics, MockResponse{
			SystemStatus: status,
			Env:          env,
			ID:           i + 1,
		})

	}
	decoder := json.NewEncoder(w)
	if err := decoder.Encode(&metrics); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println("200 /metrics")

}
