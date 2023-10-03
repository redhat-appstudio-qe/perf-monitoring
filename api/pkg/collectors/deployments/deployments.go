package deployments

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/redhat-appstudio-qe/perf-monitoring/api/pkg/constants"
)

type PerfDeploymentsCollector struct {
	SuccessfulDeploymentsCreationCounter prometheus.Counter
	FailedDeploymentsCreationCounter  prometheus.Counter
	DeploymentsCreationTimeGauge        prometheus.Gauge
	ActualDeploymentsCreationTimeGauge  prometheus.Gauge
}

var (
	SuccessfulDeploymentsCreationCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "Loadtests",
		Name:      "loadtest_successful_deployments_creations",
		Help:      "this is the total no of successful Deployments created during this test ",
	})
	FailedDeploymentsCreationCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "Loadtests",
		Name:      "loadtest_failed_deployments_creations",
		Help:      "this is the total no of failed Deployments created during this test ",
	})
	DeploymentsCreationTimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "loadtest_deployments_creation_time",
		Help:      "Deployments creation time Achieved",
	})
	ActualDeploymentsCreationTimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "loadtest_actual_deployments_creation_time",
		Help:      "Actual Deployments creation time Achieved",
	})
)

func NewPerfDeploymentsCollector() *PerfDeploymentsCollector {
	return &PerfDeploymentsCollector{
		SuccessfulDeploymentsCreationCounter: SuccessfulDeploymentsCreationCounter,
		FailedDeploymentsCreationCounter: FailedDeploymentsCreationCounter,
		DeploymentsCreationTimeGauge: DeploymentsCreationTimeGauge,
		ActualDeploymentsCreationTimeGauge: ActualDeploymentsCreationTimeGauge,
	}
}

func (P *PerfDeploymentsCollector) GetCollectors() []prometheus.Collector {
	var result []prometheus.Collector
	result = append(result, P.FailedDeploymentsCreationCounter,
		 P.SuccessfulDeploymentsCreationCounter,
		 P.DeploymentsCreationTimeGauge,
		 P.ActualDeploymentsCreationTimeGauge)
	return result
}

func (P *PerfDeploymentsCollector) IncSuccessfulDeploymentsCreationCounter(){
	P.SuccessfulDeploymentsCreationCounter.Inc()
}

func (P *PerfDeploymentsCollector) IncFailedDeploymentsCreationCounter(){
	P.FailedDeploymentsCreationCounter.Inc()
}

func (P *PerfDeploymentsCollector) SetDeploymentsCreationTimeGauge(value float64){
	P.DeploymentsCreationTimeGauge.Set(value)
}
func (P *PerfDeploymentsCollector) SetActualDeploymentsCreationTimeGauge(value float64){
	P.ActualDeploymentsCreationTimeGauge.Set(value)
}

func (P *PerfDeploymentsCollector) DecideAndPush(metricType string, metric string, values ...float64){
	if metricType == constants.MetricTypeGuage {
		P.PushGuageMetric(metric, values[0])
	}else {
		P.PushCounterMetric(metric)
	}
}

func (P *PerfDeploymentsCollector) PushCounterMetric(metric string) {
	switch metric {
	case constants.MetricSuccessfulDeploymentsCreationCounter:
		P.IncSuccessfulDeploymentsCreationCounter()
	case constants.MetricFailedDeploymentsCreationCounter:
		P.IncFailedDeploymentsCreationCounter()
	default:
		log.Fatalf(`No metric named %s found`, metric)
	}
}

func (P *PerfDeploymentsCollector) PushGuageMetric(metric string, value float64) {
	switch metric {
	case constants.MetricDeploymentsCreationTimeGauge:
		P.SetDeploymentsCreationTimeGauge(value)
	case constants.MetricActualDeploymentsCreationTimeGauge:
		P.SetActualDeploymentsCreationTimeGauge(value)
	default:
		log.Fatalf(`No metric named %s found`, metric)
	}
}

func (P *PerfDeploymentsCollector) Reset(){
	P.DeploymentsCreationTimeGauge.Set(0)
	P.ActualDeploymentsCreationTimeGauge.Set(0)
}