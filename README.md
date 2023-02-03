
# Performance Monitoring 

[![Container build & push](https://github.com/redhat-appstudio-qe/perf-monitoring/actions/workflows/build.yaml/badge.svg?branch=main)](https://github.com/redhat-appstudio-qe/perf-monitoring/actions/workflows/build.yaml)

This is a monitoring setup for when we run Load/Performance Tests on AppStudio

Runs Prometheus, Pushgateway Grafana and Ingester which communicates with each other to gather and display metrics which are captured during StoneSoup Performance/Load Tests 

Depends upon [concurency-controller](https://github.com/redhat-appstudio-qe/concurency-controller)






## The Ingester

Ingester is an API which listens for the imcomming requests which is made by the concurency-controller 
it has different paths for different metrics the concurency controller sends metrics via http `POST` method then the Ingester takes that metrics and sends it to the push gateway.

### API Reference

#### POST  Update Metrics

```http
  POST /pushMetrics
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `total` | `float64` | **Required**. Total Number of requests |
| `failed` | `float64` | **Required**. Total Number of failed requests |
| `latency` | `float64` | **Required**. Average Latency or Response Time |
| `RPS` | `float64` | **Required**. Current RPS |



## How it works (flow)

``` Concurent Controller / performance toolkit sends metrics to Ingester -->  Ingester --> Push Gateway --> Prometheus Scrape's from Push Gateway --> Grafana scrapes from the datasource prometheus ```

## Deploy on Openshift

Make sure you have openshift client pointing to the cluster where you want to deploy 
```bash
    ./deploy.sh
```

## Run Locally

Clone the project

```bash
  git clone https://github.com/redhat-appstudio-qe/perf-monitoring
```

Go to the project directory

```bash
  cd perf-monitoring
```

Run Using Docker Compose

```bash
  docker compose up
```

Rebuild 

```bash
  docker compose up --build
```

## Ports Exposed

The different ports on which the different services run 

- `Prometheus` : `9090`
- `Push Gateway` : `9091`
-  `Ingester` : `8000`
- `graphana` : `3000`

