package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rhtap-perf-test/ingester-api/pkg/metrics"
	"github.com/rhtap-perf-test/ingester-api/pkg/constants"
	"github.com/rhtap-perf-test/ingester-api/pkg/routes"
	
	
)

func main() {
	router := mux.NewRouter()

	metricsController := metrics.NewMetricController(constants.GetEnv(constants.PushgatewayUrl), constants.GetEnv(constants.JobName))
	metricsController.InitPusher()

	// API routes
	router.PathPrefix("/").HandlerFunc(routes.Home).Methods("GET")
	router.Path("/pushMetrics").HandlerFunc(routes.UpdateMetrics(metricsController)).Methods("POST")

	log.Println("Serving requests on port 8000")
	err := http.ListenAndServe(":8000", router)
	log.Fatal(err)
}