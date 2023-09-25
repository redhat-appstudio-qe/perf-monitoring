package types

type Metrics struct {
	SuccessfulUserCreationsGauge                  float64
	FailedUserCreationsGauge                      float64
	UserCreationTimeGauge                         float64
	SuccessfulApplicationCreationGauge            float64
	FailedApplicationCreationGauge                float64
	ApplicationCreationTimeGauge                  float64
	ActualApplicationCreationTimeGauge            float64
	SuccessfulCDQCreationGauge                    float64
	FailedCDQCreationGauge                        float64
	CDQCreationTimeGauge                          float64
	ActualCDQCreationTimeGauge                    float64
	SuccessfulComponentCreationGauge              float64
	FailedComponentCreationGauge                  float64
	ComponentCreationTimeGauge                    float64
	ActualComponentCreationTimeGauge              float64
	SuccessfulIntegrationTestSenarioCreationGauge float64
	FailedIntegrationTestSenarioCreationGauge     float64
	ComponentIntegrationTestSenarioTimeGauge      float64
	ActualIntegrationTestSenarioCreationTimeGauge float64
	SuccessfulPipelineRunsGauge                   float64
	FailedPipelineRunsGauge                       float64
	PipelineRunsTimeGauge                         float64
	ActualPipelineRunsTimeGauge                   float64
	SuccessfulIntegrationPipelineRunsGauge        float64
	FailedIntegrationPipelineRunsGauge            float64
	IntegrationPipelineRunsTimeGauge              float64
	ActualIntegrationPipelineRunsTimeGauge        float64
	SuccessfulDeploymentsGauge                    float64
	FailedDeploymentsGauge                        float64
	DeploymentsTimeGauge                          float64
	ActualDeploymentsTimeGauge                    float64
}
