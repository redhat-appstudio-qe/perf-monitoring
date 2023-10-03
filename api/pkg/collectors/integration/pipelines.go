package integration

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/redhat-appstudio-qe/perf-monitoring/api/pkg/constants"
)

type PerfIntegrationPipelineRunsCollector struct {
	SuccessfulIntegrationPipelineRunsCreationCounter prometheus.Counter
	FailedIntegrationPipelineRunsCreationCounter  prometheus.Counter
	IntegrationPipelineRunsCreationTimeGauge        prometheus.Gauge
	ActualIntegrationPipelineRunsCreationTimeGauge  prometheus.Gauge
	IntegrationPipelineRunsTimeGauge				 prometheus.Gauge
	ActualIntegrationPipelineRunsTimeGauge			 prometheus.Gauge
}

var (
	SuccessfulIntegrationPipelineRunsCreationCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "Loadtests",
		Name:      "loadtest_successful_IntegrationPipelineRuns_creations",
		Help:      "this is the total no of successful IntegrationPipelineRuns created during this test ",
	})
	FailedIntegrationPipelineRunsCreationCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "Loadtests",
		Name:      "loadtest_failed_IntegrationPipelineRuns_creations",
		Help:      "this is the total no of failed IntegrationPipelineRuns created during this test ",
	})
	IntegrationPipelineRunsCreationTimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "loadtest_IntegrationPipelineRuns_creation_time",
		Help:      "IntegrationPipelineRuns creation time",
	})
	ActualIntegrationPipelineRunsCreationTimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "loadtest_actual_IntegrationPipelineRuns_creation_time",
		Help:      "Actual IntegrationPipelineRuns creation time ",
	})
	IntegrationPipelineRunsTimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "loadtest_IntegrationPipelineRun_time",
		Help:      "IntegrationPipelineRuns time",
	})
	ActualIntegrationPipelineRunsTimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "loadtest_actual_IntegrationPipelineRuns_time",
		Help:      "Actual IntegrationPipelineRuns time ",
	})
)

func NewPerfIntegrationPipelineRunsCollector() *PerfIntegrationPipelineRunsCollector {
	return &PerfIntegrationPipelineRunsCollector{
		SuccessfulIntegrationPipelineRunsCreationCounter: SuccessfulIntegrationPipelineRunsCreationCounter,
		FailedIntegrationPipelineRunsCreationCounter: FailedIntegrationPipelineRunsCreationCounter,
		IntegrationPipelineRunsCreationTimeGauge: IntegrationPipelineRunsCreationTimeGauge,
		ActualIntegrationPipelineRunsCreationTimeGauge: ActualIntegrationPipelineRunsCreationTimeGauge,
		IntegrationPipelineRunsTimeGauge: IntegrationPipelineRunsTimeGauge,
		ActualIntegrationPipelineRunsTimeGauge: ActualIntegrationPipelineRunsTimeGauge,
	}
}

func (P *PerfIntegrationPipelineRunsCollector) GetCollectors() []prometheus.Collector {
	var result []prometheus.Collector
	result = append(result, P.FailedIntegrationPipelineRunsCreationCounter,
		 P.SuccessfulIntegrationPipelineRunsCreationCounter,
		 P.IntegrationPipelineRunsCreationTimeGauge,
		 P.ActualIntegrationPipelineRunsCreationTimeGauge,
		 P.IntegrationPipelineRunsTimeGauge,
		 P.ActualIntegrationPipelineRunsTimeGauge)
	return result
}

func (P *PerfIntegrationPipelineRunsCollector) IncSuccessfulIntegrationPipelineRunsCreationCounter(){
	P.SuccessfulIntegrationPipelineRunsCreationCounter.Inc()
}

func (P *PerfIntegrationPipelineRunsCollector) IncFailedIntegrationPipelineRunsCreationCounter(){
	P.FailedIntegrationPipelineRunsCreationCounter.Inc()
}

func (P *PerfIntegrationPipelineRunsCollector) SetIntegrationPipelineRunsCreationTimeGauge(value float64){
	P.IntegrationPipelineRunsCreationTimeGauge.Set(value)
}
func (P *PerfIntegrationPipelineRunsCollector) SetActualIntegrationPipelineRunsCreationTimeGauge(value float64){
	P.ActualIntegrationPipelineRunsCreationTimeGauge.Set(value)
}

func (P *PerfIntegrationPipelineRunsCollector) SetIntegrationPipelineRunsTimeGauge(value float64){
	P.IntegrationPipelineRunsTimeGauge.Set(value)
}
func (P *PerfIntegrationPipelineRunsCollector) SetActualIntegrationPipelineRunsTimeGauge(value float64){
	P.ActualIntegrationPipelineRunsTimeGauge.Set(value)
}

func (P *PerfIntegrationPipelineRunsCollector) DecideAndPush(metricType string, metric string, values ...float64){
	if metricType == constants.MetricTypeGuage {
		P.PushGuageMetric(metric, values[0])
	}else {
		P.PushCounterMetric(metric)
	}
}

func (P *PerfIntegrationPipelineRunsCollector) PushCounterMetric(metric string) {
	switch metric {
	case constants.MetricSuccessfulIntegrationPipelineRunsCreationCounter:
		P.IncSuccessfulIntegrationPipelineRunsCreationCounter()
	case constants.MetricFailedIntegrationPipelineRunsCreationCounter:
		P.IncFailedIntegrationPipelineRunsCreationCounter()
	default:
		log.Fatalf(`No metric named %s found`, metric)
	}
}

func (P *PerfIntegrationPipelineRunsCollector) PushGuageMetric(metric string, value float64) {
	switch metric {
	case constants.MetricIntegrationPipelineRunsCreationTimeGauge:
		P.SetIntegrationPipelineRunsCreationTimeGauge(value)
	case constants.MetricActualIntegrationPipelineRunsCreationTimeGauge:
		P.SetActualIntegrationPipelineRunsCreationTimeGauge(value)
	case constants.MetricIntegrationPipelineRunsTimeGauge:
		P.SetIntegrationPipelineRunsTimeGauge(value)
	case constants.MetricActualIntegrationPipelineRunsTimeGauge:
		P.SetActualIntegrationPipelineRunsTimeGauge(value)
	default:
		log.Fatalf(`No metric named %s found`, metric)
	}
}

func (P *PerfIntegrationPipelineRunsCollector) Reset(){
	P.IntegrationPipelineRunsCreationTimeGauge.Set(0)
	P.ActualIntegrationPipelineRunsCreationTimeGauge.Set(0)
	P.ActualIntegrationPipelineRunsTimeGauge.Set(0)
	P.IntegrationPipelineRunsTimeGauge.Set(0)
}