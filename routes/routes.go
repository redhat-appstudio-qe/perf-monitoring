package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/redhat-appstudio-qe/perf-monitoring/metrics"
)


type BatchWiseMetrics struct {
	T float64 `json:"total"`
	F  float64    `json:"failed"`
	S float64 `json:"success"`
}

type TotalUpdates struct {
	T float64 `json:"totalReq"`
}

type TimeMetric struct {
	T float64 `json:"time"`
}


func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello welcome "))
}

func AddBatchWise(metricsController *metrics.MetricsPush) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		var BatchWiseMetric BatchWiseMetrics
		err := json.NewDecoder(r.Body).Decode(&BatchWiseMetric)
		if err != nil {
				log.Fatalln("There was an error decoding the request body into the struct")
		}
		metricsController.IncreaseGuage(BatchWiseMetric.T, BatchWiseMetric.S, BatchWiseMetric.F)
		err = json.NewEncoder(w).Encode(&BatchWiseMetric)
		if err != nil {
				log.Fatalln("There was an error encoding the initialized struct")
		}
	}
	
}

func UpdateTotal(metricsController *metrics.MetricsPush) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        var TotalUpdate TotalUpdates
        err := json.NewDecoder(r.Body).Decode(&TotalUpdate)
        if err != nil {
                log.Fatalln("There was an error decoding the request body into the struct")
        }
        metricsController.IncreaseCountReq(TotalUpdate.T)
        err = json.NewEncoder(w).Encode(&TotalUpdate)
        if err != nil {
                log.Fatalln("There was an error encoding the initialized struct")
        }
		
	}
	
}

func UpdateAvgTime(metricsController *metrics.MetricsPush) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        var TotalUpdate TimeMetric
        err := json.NewDecoder(r.Body).Decode(&TotalUpdate)
        if err != nil {
                log.Fatalln("There was an error decoding the request body into the struct")
        }
        metricsController.AddAverageTimeReq(TotalUpdate.T)
        err = json.NewEncoder(w).Encode(&TotalUpdate)
        if err != nil {
                log.Fatalln("There was an error encoding the initialized struct")
        }
		
	}
	
}

func UpdateTime(metricsController *metrics.MetricsPush) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        var TotalUpdate TimeMetric
        err := json.NewDecoder(r.Body).Decode(&TotalUpdate)
		log.Println(TotalUpdate)
        if err != nil {
                log.Fatalln("There was an error decoding the request body into the struct")
        }
        metricsController.AddTotalTimeReq(TotalUpdate.T)
        err = json.NewEncoder(w).Encode(&TotalUpdate)
        if err != nil {
                log.Fatalln("There was an error encoding the initialized struct")
        }
		
	}
	
}