package integration

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/redhat-appstudio-qe/perf-monitoring/api/pkg/constants"
)

type PerfIntegrationTestSenarioCollector struct {
	SuccessfulIntegrationTestSenarioCreationCounter prometheus.Counter
	FailedIntegrationTestSenarioCreationCounter  prometheus.Counter
	IntegrationTestSenarioCreationTimeGauge        prometheus.Gauge
	ActualIntegrationTestSenarioCreationTimeGauge  prometheus.Gauge
}

var (
	SuccessfulIntegrationTestSenarioCreationCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "Loadtests",
		Name:      "loadtest_successful_IntegrationTestSenario_creations",
		Help:      "this is the total no of successful IntegrationTestSenario created during this test ",
	})
	FailedIntegrationTestSenarioCreationCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "Loadtests",
		Name:      "loadtest_failed_IntegrationTestSenario_creations",
		Help:      "this is the total no of failed IntegrationTestSenario created during this test ",
	})
	IntegrationTestSenarioCreationTimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "loadtest_IntegrationTestSenario_creation_time",
		Help:      "IntegrationTestSenario creation time Achieved",
	})
	ActualIntegrationTestSenarioCreationTimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "loadtest_actual_IntegrationTestSenario_creation_time",
		Help:      "Actual IntegrationTestSenario creation time Achieved",
	})
)

func NewPerfIntegrationTestSenarioCollector() *PerfIntegrationTestSenarioCollector {
	return &PerfIntegrationTestSenarioCollector{
		SuccessfulIntegrationTestSenarioCreationCounter: SuccessfulIntegrationTestSenarioCreationCounter,
		FailedIntegrationTestSenarioCreationCounter: FailedIntegrationTestSenarioCreationCounter,
		IntegrationTestSenarioCreationTimeGauge: IntegrationTestSenarioCreationTimeGauge,
		ActualIntegrationTestSenarioCreationTimeGauge: ActualIntegrationTestSenarioCreationTimeGauge,
	}
}

func (P *PerfIntegrationTestSenarioCollector) GetCollectors() []prometheus.Collector {
	var result []prometheus.Collector
	result = append(result, P.FailedIntegrationTestSenarioCreationCounter,
		 P.SuccessfulIntegrationTestSenarioCreationCounter,
		 P.IntegrationTestSenarioCreationTimeGauge,
		 P.ActualIntegrationTestSenarioCreationTimeGauge)
	return result
}

func (P *PerfIntegrationTestSenarioCollector) IncSuccessfulIntegrationTestSenarioCreationCounter(){
	P.SuccessfulIntegrationTestSenarioCreationCounter.Inc()
}

func (P *PerfIntegrationTestSenarioCollector) IncFailedIntegrationTestSenarioCreationCounter(){
	P.FailedIntegrationTestSenarioCreationCounter.Inc()
}

func (P *PerfIntegrationTestSenarioCollector) SetIntegrationTestSenarioCreationTimeGauge(value float64){
	P.IntegrationTestSenarioCreationTimeGauge.Set(value)
}
func (P *PerfIntegrationTestSenarioCollector) SetActualIntegrationTestSenarioCreationTimeGauge(value float64){
	P.ActualIntegrationTestSenarioCreationTimeGauge.Set(value)
}

func (P *PerfIntegrationTestSenarioCollector) DecideAndPush(metricType string, metric string, values ...float64){
	if metricType == constants.MetricTypeGuage {
		P.PushGuageMetric(metric, values[0])
	}else {
		P.PushCounterMetric(metric)
	}
}

func (P *PerfIntegrationTestSenarioCollector) PushCounterMetric(metric string) {
	switch metric {
	case constants.MetricSuccessfulIntegrationTestSenarioCreationCounter:
		P.IncSuccessfulIntegrationTestSenarioCreationCounter()
	case constants.MetricFailedIntegrationTestSenarioCreationCounter:
		P.IncFailedIntegrationTestSenarioCreationCounter()
	default:
		log.Fatalf(`No metric named %s found`, metric)
	}
}

func (P *PerfIntegrationTestSenarioCollector) PushGuageMetric(metric string, value float64) {
	switch metric {
	case constants.MetricIntegrationTestSenarioCreationTimeGauge:
		P.SetIntegrationTestSenarioCreationTimeGauge(value)
	case constants.MetricActualIntegrationTestSenarioCreationTimeGauge:
		P.SetActualIntegrationTestSenarioCreationTimeGauge(value)
	default:
		log.Fatalf(`No metric named %s found`, metric)
	}
}

func (P *PerfIntegrationTestSenarioCollector) Reset(){
	P.IntegrationTestSenarioCreationTimeGauge.Set(0)
	P.ActualIntegrationTestSenarioCreationTimeGauge.Set(0)
}