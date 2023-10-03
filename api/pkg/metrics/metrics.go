package metrics

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"github.com/redhat-appstudio-qe/perf-monitoring/api/pkg/collectors/applications"
	"github.com/redhat-appstudio-qe/perf-monitoring/api/pkg/collectors/cdq"
	"github.com/redhat-appstudio-qe/perf-monitoring/api/pkg/collectors/components"
	"github.com/redhat-appstudio-qe/perf-monitoring/api/pkg/collectors/deployments"
	"github.com/redhat-appstudio-qe/perf-monitoring/api/pkg/collectors/integration"
	"github.com/redhat-appstudio-qe/perf-monitoring/api/pkg/collectors/pipelines"
	"github.com/redhat-appstudio-qe/perf-monitoring/api/pkg/collectors/users"
	"github.com/redhat-appstudio-qe/perf-monitoring/api/pkg/constants"
	"github.com/redhat-appstudio-qe/perf-monitoring/api/pkg/utils"
)

type MetricsPush struct {
	PushgatewayUrl string
	Pusher         *push.Pusher
	Temp           float64
	JobName        string
}

var (
	registry *prometheus.Registry
	usersCollector *users.PerfUserCollector
	applicationsCollector   *applications.PerfApplicationsCollector
	cdqCollector *cdq.PerfCDQCollector
	componentsCollector *components.PerfComponentCollector
	deploymentsCollector *deployments.PerfDeploymentsCollector
	integrationtestscCollector *integration.PerfIntegrationTestSenarioCollector
	integrationtestpipelinesCollector *integration.PerfIntegrationPipelineRunsCollector
	pipelinesCollector *pipelines.PerfPipelineRunsCollector

)



func NewMetricController(PushgatewayUrl string, JobName string) *MetricsPush {
	registry = prometheus.NewRegistry()
	usersCollector = users.NewPerfUserCollector()
	applicationsCollector = applications.NewPerfApplicationsCollector()
	cdqCollector = cdq.NewPerfCDQCollector()
	componentsCollector = components.NewPerfComponentCollector()
	deploymentsCollector = deployments.NewPerfDeploymentsCollector()
	integrationtestscCollector = integration.NewPerfIntegrationTestSenarioCollector()
	integrationtestpipelinesCollector = integration.NewPerfIntegrationPipelineRunsCollector()
	pipelinesCollector = pipelines.NewPerfPipelineRunsCollector()
	registry.MustRegister(
		utils.CombineCollectors(
			usersCollector.GetCollectors(), 
			applicationsCollector.GetCollectors(), 
			cdqCollector.GetCollectors(),
			componentsCollector.GetCollectors(),
			deploymentsCollector.GetCollectors(),
			integrationtestscCollector.GetCollectors(),
			integrationtestpipelinesCollector.GetCollectors(),
			pipelinesCollector.GetCollectors(),
			)...
	)

	return &MetricsPush{PushgatewayUrl: PushgatewayUrl, JobName: JobName, Temp: 0.0}
}

func (M *MetricsPush) InitPusher() {
	M.Pusher = push.New(M.PushgatewayUrl, M.JobName).Gatherer(registry)
}

func pushMetric(M *MetricsPush) {
	if err := M.Pusher.Add(); err != nil {
		log.Println("Could not push to Pushgateway:", err)
	}
}

func (M *MetricsPush) PushMetrics(collector string, metricType string, metric string, values ...float64) {
	// Push metrics
	switch collector{
	case constants.CollectorUsers:
		usersCollector.DecideAndPush(metricType, metric, values...)
	case constants.CollectorApplications:
		applicationsCollector.DecideAndPush(metricType, metric, values...)
	case constants.CollectorCDQ:
		cdqCollector.DecideAndPush(metricType, metric, values...)
	case constants.CollectorComponents:
		componentsCollector.DecideAndPush(metricType, metric, values...)
	case constants.CollectorIntegrationTestsSC:
		integrationtestscCollector.DecideAndPush(metricType, metric, values...)
	case constants.CollectorIntegrationTestsPipeline:
		integrationtestpipelinesCollector.DecideAndPush(metricType, metric, values...)
	case constants.CollectorPipelines:
		pipelinesCollector.DecideAndPush(metricType, metric, values...)
	case constants.CollectorDeployments:
		deploymentsCollector.DecideAndPush(metricType, metric, values...)
	default:
		log.Fatalf("Collector %s Not Found", collector)
	}
	pushMetric(M)
}


func (M *MetricsPush) ResetMetrics() {
	usersCollector.Reset()
	applicationsCollector.Reset()
	cdqCollector.Reset()
	componentsCollector.Reset()
	integrationtestscCollector.Reset()
	integrationtestpipelinesCollector.Reset()
	pipelinesCollector.Reset()
	deploymentsCollector.Reset()
	pushMetric(M)
}

