package main

import (
	"time"

	"github.com/redhat-appstudio-qe/perf-monitoring/api/pkg/constants"
	"github.com/redhat-appstudio-qe/perf-monitoring/api/pkg/metrics"
)

func main(){
		controller := metrics.NewMetricController("<>", "load")
		controller.InitPusher();

		time.Sleep(time.Second * 2)

		for i:=0; i<10; i++ {
			controller.PushMetrics(constants.CollectorUsers, constants.MetricTypeCounter, constants.MetricSuccessfulUserCreationsCounter)
			controller.PushMetrics(constants.CollectorUsers, constants.MetricTypeGuage, constants.MetricUserCreationTimeGauge, 15.0)
			time.Sleep(time.Second * 5)
		}
}