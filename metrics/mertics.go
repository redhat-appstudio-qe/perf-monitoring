package metrics

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

var (
	registry *prometheus.Registry
	
	TotalReqGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name: "total_requests_made",
		Help: "this is the total requests made during this test ",
	})
	SuccessfulReqGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name: "total_successful_requests_made",
		Help: "this is the total no of successful requests made during this test ",
	})
	FailureReqGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name: "total_failed_requests_made",
		Help: "this is the total no of failed requests made during this test ",
	})
	TotalRequestsEntirelyGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name: "total_overall_requests_made",
		Help: "this is the total no of requests made ",
	})
	LatencyGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name: "avg_latency_achieved",
		Help: "Average latency Achieved",
	})

	RPSGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name: "current_rps",
		Help: "Current RPS acheived",
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
	registry.MustRegister(TotalReqGauge, SuccessfulReqGauge, FailureReqGauge, TotalRequestsEntirelyGauge, LatencyGauge, RPSGauge)
	
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

func (M *MetricsPush) IncreaseCountReq(T float64){
	TotalRequestsEntirelyGauge.Add(T)
	pushMetric(M)
}

func (M *MetricsPush)PushMetrics(total float64, failed float64, latency float64, RPS float64){
	TotalReqGauge.Set(total)
	TotalRequestsEntirelyGauge.Add(total)
	FailureReqGauge.Set(failed)
	SuccessfulReqGauge.Set(total - failed)
	LatencyGauge.Set(latency)
	RPSGauge.Set(RPS)
	pushMetric(M)
}