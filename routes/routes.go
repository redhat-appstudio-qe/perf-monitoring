package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/redhat-appstudio-qe/perf-monitoring/metrics"
)

type UpdateMetric struct {
	Total float64 `json:"total"`
	Failed float64 `json:"failed"`
	Latency float64 `json:"latency"`
	RPS float64 `json:"RPS"`
}

func Home(w http.ResponseWriter, r *http.Request) {
	log.Println("GET /")
	w.Write([]byte("You Are Connected to Ingester"))
}

func UpdateMetrics(metricsController *metrics.MetricsPush) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		var UpdateMetrics UpdateMetric
		err := json.NewDecoder(r.Body).Decode(&UpdateMetrics)
		if err != nil {
				log.Fatalln("There was an error decoding the request body into the struct")
		}
		log.Println("POST /pushMetrics Update Request Recieved: ", UpdateMetrics)
		metricsController.PushMetrics(UpdateMetrics.Total, 
		UpdateMetrics.Failed, UpdateMetrics.Latency,
		UpdateMetrics.RPS)
		err = json.NewEncoder(w).Encode(&UpdateMetrics)
		if err != nil {
				log.Fatalln("There was an error encoding the initialized struct")
		}
	}
	
}