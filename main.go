package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/redhat-appstudio-qe/perf-monitoring/constants"
	"github.com/redhat-appstudio-qe/perf-monitoring/metrics"
	"github.com/redhat-appstudio-qe/perf-monitoring/routes"
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