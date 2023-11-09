# Performance Monitoring 


This is a  Push/Action monitoring setup for when we run Load/Performance Tests on AppStudio

Runs Prometheus, Pushgateway Grafana, and Contains a gopkg which communicates with each other to gather and display metrics that are captured during StoneSoup Performance/Load Tests 

This application consists of 3 components

- Push gateway
- Prometheus
- Grafana

# Deployment

For this application, the above 3 components have to be deployed 

## Pushgateway

The pre-configured [yaml](https://github.com/redhat-appstudio-qe/perf-monitoring/blob/main/pushgateway/deploy/base/deployment.yaml "yaml file") file

To deploy this component in RHTAP, Create a project and add Pushgateway as a component. It’ll use the configuration mentioned earlier to deploy. You can edit the configuration file, if and when necessary.

## Prometheus

Before deploying Prometheus, we should make a config map. Here is a sample [config map](https://github.com/redhat-appstudio-qe/perf-monitoring/blob/main/prometheus/deploy/base/deployment.yaml#L1-L14 "config map") reference. Edit the file accordingly.

Then, use this [deployment yaml](https://github.com/redhat-appstudio-qe/perf-monitoring/blob/main/prometheus/deploy/base/deployment.yaml "deployment yaml") to deploy it in RHTAP 


## Grafana 

To deploy Grafana in RHTAP, here are the necessary steps 
- Create secrets 
- Configure those variables in this [file](https://github.com/redhat-appstudio-qe/perf-monitoring/blob/main/grafana/deploy/base/secret.yaml "file")

The below fields are mandatory 

`GF_DATABASE_URL`  -  URL to any Database (like RDS)

`GF_DATABASE_TYPE`  -  Provide the type DB (like POSTGRES)

`GF_SECURITY_ADMIN_PASSWORD` - Password of that DB

`GF_SECURITY_ADMIN_USER` - Username of that DB

`GF_INSTALL_PLUGINS` is an optional one

##### Next steps 
- Get the offline token
- Get access token
- `oc login`
-  `cd` to  [secret.yaml](https://github.com/redhat-appstudio-qe/perf-monitoring/blob/main/grafana/deploy/base/secret.yaml "secret.yaml")
- Run `oc apply`

These commands will help in configuring the secret 

##### After configuring the secrets 
Import grafana component to RHTAP

# Upload metrics data to this application 

Once this application is deployed, the next step is to upload the data to push the gateway

- Initialize an object for `NewMetricController` with the `Pushgateway URL` 

	##### Reference code 

		controller := metrics.NewMetricController("URL", "load")
		controller.InitPusher();


- Use the `PushMetrics` function to upload the data.

	##### Reference code 

	`controller.PushMetrics(constants.CollectorUsers,constants.MetricTypeCounter,constants.MetricSuccessfulUserCreationsCounter)`


1. Here is a [tester code](https://github.com/redhat-appstudio-qe/perf-monitoring/blob/main/api/pkg/tester/main.go "tester code")
2. Use these code blocks as a reference for uploading data to push the gateway
