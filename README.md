
# Performance Monitoring 

[![Go](https://github.com/redhat-appstudio-qe/concurency-controller/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/redhat-appstudio-qe/concurency-controller/actions/workflows/go.yml)

This is a monitoring setup for when we run Load/Performance Tests on AppStudio

Runs Prometheus, Pushgateway Grafana and Exporter which communicates with each other to gather and display metrics which are captured during StoneSoup Performance/Load Tests 

Depends upon [concurency-controller](https://github.com/redhat-appstudio-qe/concurency-controller)






## The Exporter

Exporter is an API which listens for the imcomming requests which is made by the concurency-controller 
it has different paths for different metrics the concurency controller sends metrics via http `POST` method then the exporter takes that metrics and sends it to the push gateway.

### API Reference

#### POST  add Batch Wise

```http
  POST /addBatchWise
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `total` | `float64` | **Required**. Total Number of requests made |
| `failed` | `float64` | **Required**. Total Number of failed requests made |
| `success` | `float64` | **Required**. Total Number of successful requests made |

#### Update Time

updates the time histogram variables

```http
  POST /updateAvgTime and /updateTime
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `time`      | `float64` | **Required**. `time.duration` converted to float64 |








## How it works (flow)

``` Concurent Controller / performance toolkit sends metrics to Exporter -->  Exporter --> Push Gateway --> Prometheus Scrape's from Push Gateway --> Grafana scrapes from the datasource prometheus ```


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

## Ports Exposed

The different ports on which the different services run 

- `Prometheus` : `9090`
- `Push Gateway` : `9091`
-  `exporter` : `8000`
- `graphana` : `3000`

