package applications

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/redhat-appstudio-qe/perf-monitoring/api/pkg/constants"
)

type PerfApplicationsCollector struct {
	SuccessfulApplicationCreationCounter prometheus.Counter
	FailedApplicationCreationCounter  prometheus.Counter
	ApplicationCreationTimeGauge        prometheus.Gauge
	ActualApplicationCreationTimeGauge  prometheus.Gauge
}

var (
	SuccessfulApplicationCreationCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "Loadtests",
		Name:      "loadtest_successful_application_creations",
		Help:      "this is the total no of successful applications created during this test ",
	})
	FailedApplicationCreationCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "Loadtests",
		Name:      "loadtest_failed_application_creations",
		Help:      "this is the total no of failed applications created during this test ",
	})
	ApplicationCreationTimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "loadtest_application_creation_time",
		Help:      "Application creation time Achieved",
	})
	ActualApplicationCreationTimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "loadtest_actual_application_creation_time",
		Help:      "Actual application creation time Achieved",
	})
)

func NewPerfApplicationsCollector() *PerfApplicationsCollector {
	return &PerfApplicationsCollector{
		SuccessfulApplicationCreationCounter: SuccessfulApplicationCreationCounter,
		FailedApplicationCreationCounter: FailedApplicationCreationCounter,
		ApplicationCreationTimeGauge: ApplicationCreationTimeGauge,
		ActualApplicationCreationTimeGauge: ActualApplicationCreationTimeGauge,
	}
}

func (P *PerfApplicationsCollector) GetCollectors() []prometheus.Collector {
	var result []prometheus.Collector
	result = append(result, P.FailedApplicationCreationCounter,
		 P.SuccessfulApplicationCreationCounter,
		 P.ApplicationCreationTimeGauge,
		 P.ActualApplicationCreationTimeGauge)
	return result
}

func (P *PerfApplicationsCollector) IncSuccessfulApplicationCreationCounter(){
	P.SuccessfulApplicationCreationCounter.Inc()
}

func (P *PerfApplicationsCollector) IncFailedApplicationCreationCounter(){
	P.FailedApplicationCreationCounter.Inc()
}

func (P *PerfApplicationsCollector) SetApplicationCreationTimeGauge(value float64){
	P.ApplicationCreationTimeGauge.Set(value)
}
func (P *PerfApplicationsCollector) SetActualApplicationCreationTimeGauge(value float64){
	P.ActualApplicationCreationTimeGauge.Set(value)
}

func (P *PerfApplicationsCollector) DecideAndPush(metricType string, metric string, values ...float64){
	if metricType == constants.MetricTypeGuage {
		P.PushGuageMetric(metric, values[0])
	}else {
		P.PushCounterMetric(metric)
	}
}

func (P *PerfApplicationsCollector) PushCounterMetric(metric string) {
	switch metric {
	case constants.MetricSuccessfulApplicationCreationCounter:
		P.IncSuccessfulApplicationCreationCounter()
	case constants.MetricFailedApplicationCreationCounter:
		P.IncFailedApplicationCreationCounter()
	default:
		log.Fatalf(`No metric named %s found`, metric)
	}
}

func (P *PerfApplicationsCollector) PushGuageMetric(metric string, value float64) {
	switch metric {
	case constants.MetricApplicationCreationTimeGauge:
		P.SetApplicationCreationTimeGauge(value)
	case constants.MetricActualApplicationCreationTimeGauge:
		P.SetActualApplicationCreationTimeGauge(value)
	default:
		log.Fatalf(`No metric named %s found`, metric)
	}
}

func (P *PerfApplicationsCollector) Reset(){
	P.ApplicationCreationTimeGauge.Set(0)
	P.ActualApplicationCreationTimeGauge.Set(0)
}