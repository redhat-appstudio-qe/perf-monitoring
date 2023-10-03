package users

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/redhat-appstudio-qe/perf-monitoring/api/pkg/constants"
)

type PerfUserCollector struct {
	SuccessfulUserCreationsCounter prometheus.Counter
	FailedUserCreationsCounter     prometheus.Counter
	UserCreationTimeGauge          prometheus.Gauge
}

var (
	SuccessfulUserCreationsCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "Loadtests",
		Name:      "loadtest_successful_user_creations",
		Help:      "this is the total no of successful users created during this test ",
	})
	FailedUserCreationsCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "Loadtests",
		Name:      "loadtest_failed_user_creations",
		Help:      "this is the total no of failed users created during this test",
	})
	UserCreationTimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "loadtest_usersignup_time",
		Help:      "UserSignup Creation time Achieved",
	})
)

func NewPerfUserCollector() *PerfUserCollector {
	return &PerfUserCollector{
		SuccessfulUserCreationsCounter: SuccessfulUserCreationsCounter,
		FailedUserCreationsCounter:     FailedUserCreationsCounter,
		UserCreationTimeGauge:          UserCreationTimeGauge,
	}
}

func (P *PerfUserCollector) GetCollectors() []prometheus.Collector {
	var result []prometheus.Collector
	result = append(result, P.FailedUserCreationsCounter, P.SuccessfulUserCreationsCounter, P.UserCreationTimeGauge)
	return result
}

func (P *PerfUserCollector) IncSuccessfulUserCreationsCounter() {
	P.SuccessfulUserCreationsCounter.Inc()
}

func (P *PerfUserCollector) IncFailedUserCreationsCounter() {
	P.FailedUserCreationsCounter.Inc()
}

func (P *PerfUserCollector) SetUserCreationTimeGauge(value float64) {
	P.UserCreationTimeGauge.Set(value)
}

func (P *PerfUserCollector) DecideAndPush(metricType string, metric string, values ...float64){
	if metricType == constants.MetricTypeGuage {
		P.PushGuageMetric(metric, values[0])
	}else {
		P.PushCounterMetric(metric)
	}
}

func (P *PerfUserCollector) PushCounterMetric(metric string) {
	switch metric {
	case constants.MetricSuccessfulUserCreationsCounter:
		P.IncSuccessfulUserCreationsCounter()
	case constants.MetricFailedUserCreationsCounter:
		P.IncFailedUserCreationsCounter()
	default:
		log.Fatalf(`No metric named %s found`, metric)
	}
}

func (P *PerfUserCollector) PushGuageMetric(metric string, value float64) {
	switch metric {
	case constants.MetricUserCreationTimeGauge:
		P.SetUserCreationTimeGauge(value)
	default:
		log.Fatalf(`No metric named %s found`, metric)
	}
}

func (P *PerfUserCollector) Reset() {
	P.UserCreationTimeGauge.Set(0)
}
