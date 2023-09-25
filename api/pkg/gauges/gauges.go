package gauges

import "github.com/prometheus/client_golang/prometheus"

var (
	SuccessfulUserCreationsGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "successful_user_creations",
		Help:      "this is the total no of successful users created during this test ",
	})
	FailedUserCreationsGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "failed_user_creations",
		Help:      "this is the total no of failed users created during this test ",
	})
	UserCreationTimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "usersignup_time",
		Help:      "Average UserSignup Creation time Achieved",
	})
	SuccessfulApplicationCreationGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "successful_application_creations",
		Help:      "this is the total no of successful applications created during this test ",
	})
	FailedApplicationCreationGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "failed_application_creations",
		Help:      "this is the total no of failed applications created during this test ",
	})
	ApplicationCreationTimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "application_creation_time",
		Help:      "Application creation time Achieved",
	})
	ActualApplicationCreationTimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "actual_application_creation_time",
		Help:      "Actual application creation time Achieved",
	})
	SuccessfulCDQCreationGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "successful_cdq_creations",
		Help:      "this is the total no of successful CDQ created during this test ",
	})
	FailedCDQCreationGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "failed_cdq_creations",
		Help:      "this is the total no of failed CDQ created during this test ",
	})
	CDQCreationTimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "application_creation_time",
		Help:      "CDQ creation time Achieved",
	})
	ActualCDQCreationTimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "actual_cdq_creation_time",
		Help:      "Actual CDQ creation time Achieved",
	})
	SuccessfulComponentCreationGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "successful_component_creations",
		Help:      "this is the total no of successful Component created during this test ",
	})
	FailedComponentCreationGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "failed_component_creations",
		Help:      "this is the total no of failed Component created during this test ",
	})
	ComponentCreationTimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "component_creation_time",
		Help:      "Component creation time Achieved",
	})
	ActualComponentCreationTimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "actual_component_creation_time",
		Help:      "Actual Component creation time Achieved",
	})
	SuccessfulIntegrationTestSenarioCreationGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "successful_integrationtestsenario_creations",
		Help:      "this is the total no of successful IntegrationTestSenario created during this test ",
	})
	FailedIntegrationTestSenarioCreationGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "failed_integrationtestsenario_creations",
		Help:      "this is the total no of failed IntegrationTestSenario created during this test ",
	})
	ComponentIntegrationTestSenarioTimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "integrationtestsenario_creation_time",
		Help:      "IntegrationTestSenario creation time Achieved",
	})
	ActualIntegrationTestSenarioCreationTimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "actual_integrationtestsenario_creation_time",
		Help:      "Actual IntegrationTestSenario creation time Achieved",
	})
	SuccessfulPipelineRunsGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "successful_pipeline_runs",
		Help:      "this is the total no of successful Pipeline Runs during this test ",
	})
	FailedPipelineRunsGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "failed_pipeline_runs",
		Help:      "this is the total no of failed Pipeline Runs during this test ",
	})
	PipelineRunsTimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "pipeline_runs_time",
		Help:      "Pipeline Runs time Achieved",
	})
	ActualPipelineRunsTimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "actual_pipeline_runs_time",
		Help:      "Actual Pipeline Runs time Achieved",
	})
	SuccessfulIntegrationPipelineRunsGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "successful_integration_pipeline_runs",
		Help:      "this is the total no of successful Integration Pipeline Runs during this test ",
	})
	FailedIntegrationPipelineRunsGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "failed_integration_pipeline_runs",
		Help:      "this is the total no of failed Integration Pipeline Runs during this test ",
	})
	IntegrationPipelineRunsTimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "integration_pipeline_runs_time",
		Help:      "Integration Pipeline Runs time Achieved",
	})
	ActualIntegrationPipelineRunsTimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "actual_integration_pipeline_runs_time",
		Help:      "Actual Integration Pipeline Runs time Achieved",
	})
	SuccessfulDeploymentsGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "successful_deployments",
		Help:      "this is the total no of successful Deployments during this test ",
	})
	FailedDeploymentsGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "failed_deployments",
		Help:      "this is the total no of failed Deployments during this test ",
	})
	DeploymentsTimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "deployments_time",
		Help:      "Deployments time Achieved",
	})
	ActualDeploymentsTimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Loadtests",
		Name:      "actual_deployments_time",
		Help:      "Actual Deployments time Achieved",
	})
)

func Reset() {
	SuccessfulUserCreationsGauge.Set(0)
	FailedUserCreationsGauge.Set(0)
	UserCreationTimeGauge.Set(0)
	SuccessfulApplicationCreationGauge.Set(0)
	FailedApplicationCreationGauge.Set(0)
	ApplicationCreationTimeGauge.Set(0)
	ActualApplicationCreationTimeGauge.Set(0)
	SuccessfulCDQCreationGauge.Set(0)
	FailedCDQCreationGauge.Set(0)
	CDQCreationTimeGauge.Set(0)
	ActualCDQCreationTimeGauge.Set(0)
	SuccessfulComponentCreationGauge.Set(0)
	FailedComponentCreationGauge.Set(0)
	ComponentCreationTimeGauge.Set(0)
	ActualComponentCreationTimeGauge.Set(0)
	SuccessfulIntegrationTestSenarioCreationGauge.Set(0)
	FailedIntegrationTestSenarioCreationGauge.Set(0)
	ComponentIntegrationTestSenarioTimeGauge.Set(0)
	ActualIntegrationTestSenarioCreationTimeGauge.Set(0)
	SuccessfulPipelineRunsGauge.Set(0)
	FailedPipelineRunsGauge.Set(0)
	PipelineRunsTimeGauge.Set(0)
	ActualPipelineRunsTimeGauge.Set(0)
	SuccessfulIntegrationPipelineRunsGauge.Set(0)
	FailedIntegrationPipelineRunsGauge.Set(0)
	IntegrationPipelineRunsTimeGauge.Set(0)
	ActualIntegrationPipelineRunsTimeGauge.Set(0)
	SuccessfulDeploymentsGauge.Set(0)
	FailedDeploymentsGauge.Set(0)
	DeploymentsTimeGauge.Set(0)
	ActualDeploymentsTimeGauge.Set(0)

}
