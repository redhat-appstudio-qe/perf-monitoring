package components

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/redhat-appstudio-qe/perf-monitoring/api/pkg/constants"
)

type PerfComponentCollector struct {
	SuccessfulComponentCreationCounter prometheus.Counter
	FailedComponentCreationCounter  prometheus.Counter
	ComponentCreationTimeGauge        prometheus.Gauge
	ActualComponentCreationTimeGauge  prometheus.Gauge
}

var (
	SuccessfulComponentCreationCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "Loadtests",
		Name:      "loadtest_successful_component_creations",
		Help:      "this is the total no of successful Component created during this test ",
	})
	FailedComponentCreationCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "Loadtests",
		Name:      "loadtest_failed_component_creations",
		Help:      "this is the total no of failed Component created during this test ",
	})
	ComponentCreationTimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "loadtest_component_creation_time",
		Help:      "Component creation time Achieved",
	})
	ActualComponentCreationTimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "loadtest_actual_component_creation_time",
		Help:      "Actual Component creation time Achieved",
	})
)

func NewPerfComponentCollector() *PerfComponentCollector {
	return &PerfComponentCollector{
		SuccessfulComponentCreationCounter: SuccessfulComponentCreationCounter,
		FailedComponentCreationCounter: FailedComponentCreationCounter,
		ComponentCreationTimeGauge: ComponentCreationTimeGauge,
		ActualComponentCreationTimeGauge: ActualComponentCreationTimeGauge,
	}
}

func (P *PerfComponentCollector) GetCollectors() []prometheus.Collector {
	var result []prometheus.Collector
	result = append(result, P.FailedComponentCreationCounter,
		 P.SuccessfulComponentCreationCounter,
		 P.ComponentCreationTimeGauge,
		 P.ActualComponentCreationTimeGauge)
	return result
}

func (P *PerfComponentCollector) IncSuccessfulComponentCreationCounter(){
	P.SuccessfulComponentCreationCounter.Inc()
}

func (P *PerfComponentCollector) IncFailedComponentCreationCounter(){
	P.FailedComponentCreationCounter.Inc()
}

func (P *PerfComponentCollector) SetComponentCreationTimeGauge(value float64){
	P.ComponentCreationTimeGauge.Set(value)
}
func (P *PerfComponentCollector) SetActualComponentCreationTimeGauge(value float64){
	P.ActualComponentCreationTimeGauge.Set(value)
}

func (P *PerfComponentCollector) DecideAndPush(metricType string, metric string, values ...float64){
	if metricType == constants.MetricTypeGuage {
		P.PushGuageMetric(metric, values[0])
	}else {
		P.PushCounterMetric(metric)
	}
}

func (P *PerfComponentCollector) PushCounterMetric(metric string) {
	switch metric {
	case constants.MetricSuccessfulComponentCreationCounter:
		P.IncSuccessfulComponentCreationCounter()
	case constants.MetricFailedComponentCreationCounter:
		P.IncFailedComponentCreationCounter()
	default:
		log.Fatalf(`No metric named %s found`, metric)
	}
}

func (P *PerfComponentCollector) PushGuageMetric(metric string, value float64) {
	switch metric {
	case constants.MetricComponentCreationTimeGauge:
		P.SetComponentCreationTimeGauge(value)
	case constants.MetricActualComponentCreationTimeGauge:
		P.SetActualComponentCreationTimeGauge(value)
	default:
		log.Fatalf(`No metric named %s found`, metric)
	}
}

func (P *PerfComponentCollector) Reset(){
	P.ComponentCreationTimeGauge.Set(0)
	P.ActualComponentCreationTimeGauge.Set(0)
}