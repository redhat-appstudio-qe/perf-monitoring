package metrics

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"github.com/rhtap-perf-test/ingester-api/pkg/gauges"
	"github.com/rhtap-perf-test/ingester-api/pkg/types"
)

type MetricsPush struct {
	PushgatewayUrl string
	Pusher         *push.Pusher
	Temp           float64
	JobName        string
}

var (
	registry *prometheus.Registry
)

func NewMetricController(PushgatewayUrl string, JobName string) *MetricsPush {
	registry = prometheus.NewRegistry()
	registry.MustRegister(
		gauges.SuccessfulUserCreationsGauge,
		gauges.FailedUserCreationsGauge,
		gauges.UserCreationTimeGauge,
		gauges.SuccessfulApplicationCreationGauge,
		gauges.FailedApplicationCreationGauge,
		gauges.ApplicationCreationTimeGauge,
		gauges.ActualApplicationCreationTimeGauge,
		gauges.SuccessfulCDQCreationGauge,
		gauges.FailedCDQCreationGauge,
		gauges.CDQCreationTimeGauge,
		gauges.ActualCDQCreationTimeGauge,
		gauges.SuccessfulComponentCreationGauge,
		gauges.FailedComponentCreationGauge,
		gauges.ComponentCreationTimeGauge,
		gauges.ActualComponentCreationTimeGauge,
		gauges.SuccessfulIntegrationTestSenarioCreationGauge,
		gauges.FailedIntegrationTestSenarioCreationGauge,
		gauges.ComponentIntegrationTestSenarioTimeGauge,
		gauges.ActualIntegrationTestSenarioCreationTimeGauge,
		gauges.SuccessfulPipelineRunsGauge,
		gauges.FailedPipelineRunsGauge,
		gauges.PipelineRunsTimeGauge,
		gauges.ActualPipelineRunsTimeGauge,
		gauges.SuccessfulIntegrationPipelineRunsGauge,
		gauges.FailedIntegrationPipelineRunsGauge,
		gauges.IntegrationPipelineRunsTimeGauge,
		gauges.ActualIntegrationPipelineRunsTimeGauge,
		gauges.SuccessfulDeploymentsGauge,
		gauges.FailedDeploymentsGauge,
		gauges.DeploymentsTimeGauge,
		gauges.ActualDeploymentsTimeGauge,
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

func (M *MetricsPush) PushMetrics(inputMetrics types.Metrics) {

	// Check each field and store non-zero values in the new variable
	if inputMetrics.SuccessfulUserCreationsGauge != 0 {
		gauges.SuccessfulUserCreationsGauge.Set(inputMetrics.SuccessfulUserCreationsGauge)
	}
	if inputMetrics.FailedUserCreationsGauge != 0 {
		gauges.FailedUserCreationsGauge.Set(inputMetrics.FailedUserCreationsGauge)
	}
	if inputMetrics.UserCreationTimeGauge != 0 {
		gauges.UserCreationTimeGauge.Set(inputMetrics.UserCreationTimeGauge)
	}
	if inputMetrics.SuccessfulApplicationCreationGauge != 0 {
		gauges.SuccessfulApplicationCreationGauge.Set(inputMetrics.SuccessfulApplicationCreationGauge)
	}
	if inputMetrics.FailedApplicationCreationGauge != 0 {
		gauges.FailedApplicationCreationGauge.Set(inputMetrics.FailedApplicationCreationGauge)
	}
	if inputMetrics.ApplicationCreationTimeGauge != 0 {
		gauges.ApplicationCreationTimeGauge.Set(inputMetrics.ApplicationCreationTimeGauge)
	}
	if inputMetrics.ActualApplicationCreationTimeGauge != 0 {
		gauges.ActualApplicationCreationTimeGauge.Set(inputMetrics.ActualApplicationCreationTimeGauge)
	}
	if inputMetrics.SuccessfulCDQCreationGauge != 0 {
		gauges.SuccessfulCDQCreationGauge.Set(inputMetrics.SuccessfulCDQCreationGauge)
	}
	if inputMetrics.FailedCDQCreationGauge != 0 {
		gauges.FailedCDQCreationGauge.Set(inputMetrics.FailedCDQCreationGauge)
	}
	if inputMetrics.CDQCreationTimeGauge != 0 {
		gauges.CDQCreationTimeGauge.Set(inputMetrics.CDQCreationTimeGauge)
	}
	if inputMetrics.ActualCDQCreationTimeGauge != 0 {
		gauges.ActualCDQCreationTimeGauge.Set(inputMetrics.ActualCDQCreationTimeGauge)
	}
	if inputMetrics.SuccessfulComponentCreationGauge != 0 {
		gauges.SuccessfulComponentCreationGauge.Set(inputMetrics.SuccessfulComponentCreationGauge)
	}
	if inputMetrics.FailedComponentCreationGauge != 0 {
		gauges.FailedComponentCreationGauge.Set(inputMetrics.FailedComponentCreationGauge)
	}
	if inputMetrics.ComponentCreationTimeGauge != 0 {
		gauges.ComponentCreationTimeGauge.Set(inputMetrics.ComponentCreationTimeGauge)
	}
	if inputMetrics.ActualComponentCreationTimeGauge != 0 {
		gauges.ActualComponentCreationTimeGauge.Set(inputMetrics.ActualComponentCreationTimeGauge)
	}
	if inputMetrics.SuccessfulIntegrationTestSenarioCreationGauge != 0 {
		gauges.SuccessfulIntegrationTestSenarioCreationGauge.Set(inputMetrics.SuccessfulIntegrationTestSenarioCreationGauge)
	}
	if inputMetrics.FailedIntegrationTestSenarioCreationGauge != 0 {
		gauges.FailedIntegrationTestSenarioCreationGauge.Set(inputMetrics.FailedIntegrationTestSenarioCreationGauge)
	}
	if inputMetrics.ComponentIntegrationTestSenarioTimeGauge != 0 {
		gauges.ComponentIntegrationTestSenarioTimeGauge.Set(inputMetrics.ComponentIntegrationTestSenarioTimeGauge)
	}
	if inputMetrics.ActualIntegrationTestSenarioCreationTimeGauge != 0 {
		gauges.ActualIntegrationTestSenarioCreationTimeGauge.Set(inputMetrics.ActualIntegrationTestSenarioCreationTimeGauge)
	}
	if inputMetrics.SuccessfulPipelineRunsGauge != 0 {
		gauges.SuccessfulPipelineRunsGauge.Set(inputMetrics.SuccessfulPipelineRunsGauge)
	}
	if inputMetrics.FailedPipelineRunsGauge != 0 {
		gauges.FailedPipelineRunsGauge.Set(inputMetrics.FailedPipelineRunsGauge)
	}
	if inputMetrics.PipelineRunsTimeGauge != 0 {
		gauges.PipelineRunsTimeGauge.Set(inputMetrics.PipelineRunsTimeGauge)
	}
	if inputMetrics.ActualPipelineRunsTimeGauge != 0 {
		gauges.ActualPipelineRunsTimeGauge.Set(inputMetrics.ActualPipelineRunsTimeGauge)
	}
	if inputMetrics.SuccessfulIntegrationPipelineRunsGauge != 0 {
		gauges.SuccessfulIntegrationPipelineRunsGauge.Set(inputMetrics.SuccessfulIntegrationPipelineRunsGauge)
	}
	if inputMetrics.FailedIntegrationPipelineRunsGauge != 0 {
		gauges.FailedIntegrationPipelineRunsGauge.Set(inputMetrics.FailedIntegrationPipelineRunsGauge)
	}
	if inputMetrics.IntegrationPipelineRunsTimeGauge != 0 {
		gauges.IntegrationPipelineRunsTimeGauge.Set(inputMetrics.IntegrationPipelineRunsTimeGauge)
	}
	if inputMetrics.ActualIntegrationPipelineRunsTimeGauge != 0 {
		gauges.ActualIntegrationPipelineRunsTimeGauge.Set(inputMetrics.ActualIntegrationPipelineRunsTimeGauge)
	}
	if inputMetrics.SuccessfulDeploymentsGauge != 0 {
		gauges.SuccessfulDeploymentsGauge.Set(inputMetrics.SuccessfulDeploymentsGauge)
	}
	if inputMetrics.FailedDeploymentsGauge != 0 {
		gauges.FailedDeploymentsGauge.Set(inputMetrics.FailedDeploymentsGauge)
	}
	if inputMetrics.DeploymentsTimeGauge != 0 {
		gauges.DeploymentsTimeGauge.Set(inputMetrics.DeploymentsTimeGauge)
	}
	if inputMetrics.ActualDeploymentsTimeGauge != 0 {
		gauges.ActualDeploymentsTimeGauge.Set(inputMetrics.ActualDeploymentsTimeGauge)
	}

	// Push metrics
	pushMetric(M)
}

func (M *MetricsPush) ResetMetrics() {
	gauges.Reset()
	pushMetric(M)
}

