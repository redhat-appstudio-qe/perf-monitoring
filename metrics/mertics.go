package metrics

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

var (
	registry *prometheus.Registry
	
	TotalReq = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name: "total_requests_made",
		Help: "this is the total requests made during this test ",
	})
	SuccessfulReq = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name: "total_successful_requests_made",
		Help: "this is the total no of successful requests made during this test ",
	})
	FailureReq = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name: "total_failed_requests_made",
		Help: "this is the total no of failed requests made during this test ",
	})
	TotalRequestsEntirely = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name: "total_overall_requests_made",
		Help: "this is the total no of requests made ",
	})

	AVGTimeByBatch = prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: "Loadtests",
		Name: "avgtime_taken_by_batch",
		Help: "This is average time taken by Batches",
	})

	TotalTestTime = prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: "Loadtests",
		Name: "test_run_time",
		Help: "This is total time taken by test",
	})
	
)


type MetricsPush struct {
	PushgatewayUrl string
	pusher *push.Pusher
	Temp float64
	JobName string
}

func NewMetricController(PushgatewayUrl string, JobName string) *MetricsPush {
	registry = prometheus.NewRegistry()
	registry.MustRegister(TotalReq, SuccessfulReq, FailureReq, TotalRequestsEntirely, AVGTimeByBatch, TotalTestTime)
	
	return &MetricsPush{PushgatewayUrl: PushgatewayUrl, JobName: JobName, Temp: 0.0}
}

func (M *MetricsPush) InitPusher() {
	M.pusher = push.New(M.PushgatewayUrl, M.JobName).Gatherer(registry)
	
}

func pushMetric(M *MetricsPush){
	if err := M.pusher.Add(); err != nil {
		log.Println("Could not push to Pushgateway:", err)
	}
}

func (M *MetricsPush) IncreaseGuage(T float64, S float64, F float64) {
	TotalReq.Set(T)
	SuccessfulReq.Set(S)
	FailureReq.Set(F)
	if err := M.pusher.Add(); err != nil {
		log.Println("Could not push to Pushgateway:", err)
	}
}

func (M *MetricsPush) IncreaseCountReq(T float64){
	TotalRequestsEntirely.Add(T)
	pushMetric(M)
}

func (M *MetricsPush)AddAverageTimeReq(d float64){
	AVGTimeByBatch.Observe(d)
	pushMetric(M)
}

func (M *MetricsPush)AddTotalTimeReq(d float64){
	TotalTestTime.Observe(d)
	pushMetric(M)
}