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

	// Serving static files
	router.PathPrefix("/").HandlerFunc(routes.Home).Methods("GET")
	router.PathPrefix("/addBatchWise").HandlerFunc(routes.AddBatchWise(metricsController)).Methods("POST")
	router.PathPrefix("/updateTotal").HandlerFunc(routes.UpdateTotal(metricsController)).Methods("POST")
	router.PathPrefix("/updateAvgTime").HandlerFunc(routes.UpdateAvgTime(metricsController)).Methods("POST")
	router.PathPrefix("/updateTime").HandlerFunc(routes.UpdateTime(metricsController)).Methods("POST")

	log.Println("Serving requests on port 8000")
	err := http.ListenAndServe(":8000", router)
	log.Fatal(err)
}