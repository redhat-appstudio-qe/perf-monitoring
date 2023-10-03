package cdq

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/redhat-appstudio-qe/perf-monitoring/api/pkg/constants"
)

type PerfCDQCollector struct {
	SuccessfulCDQCreationCounter prometheus.Counter
	FailedCDQCreationCounter  prometheus.Counter
	CDQCreationTimeGauge        prometheus.Gauge
	ActualCDQCreationTimeGauge  prometheus.Gauge
}

var (
	SuccessfulCDQCreationCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "Loadtests",
		Name:      "loadtest_successful_CDQ_creations",
		Help:      "this is the total no of successful CDQ created during this test ",
	})
	FailedCDQCreationCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "Loadtests",
		Name:      "loadtest_failed_CDQ_creations",
		Help:      "this is the total no of failed CDQ created during this test ",
	})
	CDQCreationTimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "loadtest_CDQ_creation_time",
		Help:      "CDQ creation time Achieved",
	})
	ActualCDQCreationTimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "loadtest_actual_CDQ_creation_time",
		Help:      "Actual CDQ creation time Achieved",
	})
)

func NewPerfCDQCollector() *PerfCDQCollector {
	return &PerfCDQCollector{
		SuccessfulCDQCreationCounter: SuccessfulCDQCreationCounter,
		FailedCDQCreationCounter: FailedCDQCreationCounter,
		CDQCreationTimeGauge: CDQCreationTimeGauge,
		ActualCDQCreationTimeGauge: ActualCDQCreationTimeGauge,
	}
}

func (P *PerfCDQCollector) GetCollectors() []prometheus.Collector {
	var result []prometheus.Collector
	result = append(result, P.FailedCDQCreationCounter,
		 P.SuccessfulCDQCreationCounter,
		 P.CDQCreationTimeGauge,
		 P.ActualCDQCreationTimeGauge)
	return result
}

func (P *PerfCDQCollector) IncSuccessfulCDQCreationCounter(){
	P.SuccessfulCDQCreationCounter.Inc()
}

func (P *PerfCDQCollector) IncFailedCDQCreationCounter(){
	P.FailedCDQCreationCounter.Inc()
}

func (P *PerfCDQCollector) SetCDQCreationTimeGauge(value float64){
	P.CDQCreationTimeGauge.Set(value)
}
func (P *PerfCDQCollector) SetActualCDQCreationTimeGauge(value float64){
	P.ActualCDQCreationTimeGauge.Set(value)
}

func (P *PerfCDQCollector) DecideAndPush(metricType string, metric string, values ...float64){
	if metricType == constants.MetricTypeGuage {
		P.PushGuageMetric(metric, values[0])
	}else {
		P.PushCounterMetric(metric)
	}
}

func (P *PerfCDQCollector) PushCounterMetric(metric string) {
	switch metric {
	case constants.MetricSuccessfulCDQCreationCounter:
		P.IncSuccessfulCDQCreationCounter()
	case constants.MetricFailedCDQCreationCounter:
		P.IncFailedCDQCreationCounter()
	default:
		log.Fatalf(`No metric named %s found`, metric)
	}
}

func (P *PerfCDQCollector) PushGuageMetric(metric string, value float64) {
	switch metric {
	case constants.MetricCDQCreationTimeGauge:
		P.SetCDQCreationTimeGauge(value)
	case constants.MetricActualCDQCreationTimeGauge:
		P.SetActualCDQCreationTimeGauge(value)
	default:
		log.Fatalf(`No metric named %s found`, metric)
	}
}

func (P *PerfCDQCollector) Reset(){
	P.CDQCreationTimeGauge.Set(0)
	P.ActualCDQCreationTimeGauge.Set(0)
}