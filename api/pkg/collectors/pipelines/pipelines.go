package pipelines

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/redhat-appstudio-qe/perf-monitoring/api/pkg/constants"
)

type PerfPipelineRunsCollector struct {
	SuccessfulPipelineRunsCreationCounter prometheus.Counter
	FailedPipelineRunsCreationCounter  prometheus.Counter
	PipelineRunsCreationTimeGauge        prometheus.Gauge
	ActualPipelineRunsCreationTimeGauge  prometheus.Gauge
	PipelineRunsTimeGauge				 prometheus.Gauge
	ActualPipelineRunsTimeGauge			 prometheus.Gauge
}

var (
	SuccessfulPipelineRunsCreationCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "Loadtests",
		Name:      "loadtest_successful_pipelineruns_creations",
		Help:      "this is the total no of successful PipelineRuns created during this test ",
	})
	FailedPipelineRunsCreationCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "Loadtests",
		Name:      "loadtest_failed_pipelineruns_creations",
		Help:      "this is the total no of failed PipelineRuns created during this test ",
	})
	PipelineRunsCreationTimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "loadtest_pipelineruns_creation_time",
		Help:      "PipelineRuns creation time",
	})
	ActualPipelineRunsCreationTimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "loadtest_actual_pipelineruns_creation_time",
		Help:      "Actual PipelineRuns creation time ",
	})
	PipelineRunsTimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "loadtest_pipelinerun_time",
		Help:      "PipelineRuns time",
	})
	ActualPipelineRunsTimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "loadtest_actual_pipelineruns_time",
		Help:      "Actual PipelineRuns time ",
	})
)

func NewPerfPipelineRunsCollector() *PerfPipelineRunsCollector {
	return &PerfPipelineRunsCollector{
		SuccessfulPipelineRunsCreationCounter: SuccessfulPipelineRunsCreationCounter,
		FailedPipelineRunsCreationCounter: FailedPipelineRunsCreationCounter,
		PipelineRunsCreationTimeGauge: PipelineRunsCreationTimeGauge,
		ActualPipelineRunsCreationTimeGauge: ActualPipelineRunsCreationTimeGauge,
		PipelineRunsTimeGauge: PipelineRunsTimeGauge,
		ActualPipelineRunsTimeGauge: ActualPipelineRunsTimeGauge,
	}
}

func (P *PerfPipelineRunsCollector) GetCollectors() []prometheus.Collector {
	var result []prometheus.Collector
	result = append(result, P.FailedPipelineRunsCreationCounter,
		 P.SuccessfulPipelineRunsCreationCounter,
		 P.PipelineRunsCreationTimeGauge,
		 P.ActualPipelineRunsCreationTimeGauge,
		 P.PipelineRunsTimeGauge,
		 P.ActualPipelineRunsTimeGauge)
	return result
}

func (P *PerfPipelineRunsCollector) IncSuccessfulPipelineRunsCreationCounter(){
	P.SuccessfulPipelineRunsCreationCounter.Inc()
}

func (P *PerfPipelineRunsCollector) IncFailedPipelineRunsCreationCounter(){
	P.FailedPipelineRunsCreationCounter.Inc()
}

func (P *PerfPipelineRunsCollector) SetPipelineRunsCreationTimeGauge(value float64){
	P.PipelineRunsCreationTimeGauge.Set(value)
}
func (P *PerfPipelineRunsCollector) SetActualPipelineRunsCreationTimeGauge(value float64){
	P.ActualPipelineRunsCreationTimeGauge.Set(value)
}

func (P *PerfPipelineRunsCollector) SetPipelineRunsTimeGauge(value float64){
	P.PipelineRunsTimeGauge.Set(value)
}
func (P *PerfPipelineRunsCollector) SetActualPipelineRunsTimeGauge(value float64){
	P.ActualPipelineRunsTimeGauge.Set(value)
}

func (P *PerfPipelineRunsCollector) DecideAndPush(metricType string, metric string, values ...float64){
	if metricType == constants.MetricTypeGuage {
		P.PushGuageMetric(metric, values[0])
	}else {
		P.PushCounterMetric(metric)
	}
}

func (P *PerfPipelineRunsCollector) PushCounterMetric(metric string) {
	switch metric {
	case constants.MetricSuccessfulPipelineRunsCreationCounter:
		P.IncSuccessfulPipelineRunsCreationCounter()
	case constants.MetricFailedPipelineRunsCreationCounter:
		P.IncFailedPipelineRunsCreationCounter()
	default:
		log.Fatalf(`No metric named %s found`, metric)
	}
}

func (P *PerfPipelineRunsCollector) PushGuageMetric(metric string, value float64) {
	switch metric {
	case constants.MetricPipelineRunsCreationTimeGauge:
		P.SetPipelineRunsCreationTimeGauge(value)
	case constants.MetricActualPipelineRunsCreationTimeGauge:
		P.SetActualPipelineRunsCreationTimeGauge(value)
	case constants.MetricPipelineRunsTimeGauge:
		P.SetPipelineRunsTimeGauge(value)
	case constants.MetricActualPipelineRunsTimeGauge:
		P.SetActualPipelineRunsTimeGauge(value)
	default:
		log.Fatalf(`No metric named %s found`, metric)
	}
}

func (P *PerfPipelineRunsCollector) Reset(){
	P.PipelineRunsCreationTimeGauge.Set(0)
	P.ActualPipelineRunsCreationTimeGauge.Set(0)
	P.ActualPipelineRunsTimeGauge.Set(0)
	P.PipelineRunsTimeGauge.Set(0)
}