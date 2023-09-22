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
	UserSignupCreationTimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name: "usersignup_time",
		Help: "Average UserSignup Creation time Achieved",
	})
	ResourceCreationTimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name: "resourcecreation_time",
		Help: "Average Resource Creation time Achieved",
	})
	PipelineRunTimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name: "pipelinerun_time",
		Help: "Average PipelineRun time Achieved",
	})
	FailedUserCreationGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name: "failed_usersignups",
		Help: "Failed User Signups",
	})
	SuccessUserCreationGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name: "successful_usersignups",
		Help: "Successful User Signups",
	})
	FailedResourceCreationGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name: "failed_resourcecreation",
		Help: "Failed Resource Creations",
	})
	FailedPipelineRunsGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name: "failed_pipelineruns",
		Help: "Failed pipelineruns",
	})
	ApplicationCreationTimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name: "applicationcreation_time",
		Help: "Application Creation Gauge",
	})
	CDQCreationTimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name: "cdqcreation_time",
		Help: "CDQ Creation Gauge",
	})
	ComponentCreationTimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name: "componentcreation_time",
		Help: "Component Creation Gauge",
	})
	ActualComponentCreationTimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name: "actualcomponentcreation_time",
		Help: "Actual Component Creation Gauge",
	})
	ActualCDQCreationTimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name: "actualcdqcreation_time",
		Help: "Actual CDQ Creation Gauge",
	})
	ActualApplicationCreationTimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name: "actualapplicationcreation_time",
		Help: "Actual Application Creation Gauge",
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
	registry.MustRegister(
		TotalReqGauge, 
		SuccessfulReqGauge, 
		FailureReqGauge, 
		TotalRequestsEntirelyGauge, 
		LatencyGauge, 
		RPSGauge, 
		PipelineRunTimeGauge, 
		FailedPipelineRunsGauge, 
		FailedResourceCreationGauge, 
		ResourceCreationTimeGauge, 
		FailedUserCreationGauge,
		SuccessUserCreationGauge, 
		UserSignupCreationTimeGauge,
		ApplicationCreationTimeGauge,
		ActualApplicationCreationTimeGauge,
		CDQCreationTimeGauge,
		ActualCDQCreationTimeGauge,
		ComponentCreationTimeGauge,
		ActualComponentCreationTimeGauge,

	)
	
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

func (M *MetricsPush)PushMetricsTotal(total float64)  {
	TotalReqGauge.Set(total)
	
	pushMetric(M)
}

func (M *MetricsPush)PushMetricsResources(
	failed_resourcecreations float64, 
	latency_resourcecreation float64) {
	FailedResourceCreationGauge.Set(failed_resourcecreations)
	ResourceCreationTimeGauge.Set(latency_resourcecreation)
	pushMetric(M)
}

func (M *MetricsPush) PushApplicationMetrics(applicationCreatedAt, actualApplicationCreated float64){
	ApplicationCreationTimeGauge.Set(applicationCreatedAt)
	ActualApplicationCreationTimeGauge.Set(actualApplicationCreated)
	pushMetric(M)
}

func (M *MetricsPush) PushCDQMetrics(cdqCreatedAt, actualcdqCreated float64){
	CDQCreationTimeGauge.Set(actualcdqCreated)
	ActualCDQCreationTimeGauge.Set(actualcdqCreated)
	pushMetric(M)
}

func (M *MetricsPush) PushComponentMetrics(componentCreatedAt, actualcomponentCreated float64){
	ComponentCreationTimeGauge.Set(componentCreatedAt)
	ActualComponentCreationTimeGauge.Set(actualcomponentCreated)
	pushMetric(M)
}


func (M *MetricsPush)PushMetricsUsers(
	failed_usersignups float64,
	successful_usersignups float64,
	latency_usersignup float64) {
	FailedUserCreationGauge.Set(failed_usersignups)
	SuccessUserCreationGauge.Set(successful_usersignups)
	UserSignupCreationTimeGauge.Set(latency_usersignup)
	
	pushMetric(M)
}

func (M *MetricsPush)PushMetricsPipelines(
	failed_pipelineruns float64, 
	latency_pipelinerun float64) {
		
	FailedPipelineRunsGauge.Set(failed_pipelineruns)
	PipelineRunTimeGauge.Set(latency_pipelinerun)
	
	pushMetric(M)
}