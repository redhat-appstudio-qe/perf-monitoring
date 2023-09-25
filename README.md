
# Performance Monitoring 

[![Container build & push](https://github.com/redhat-appstudio-qe/perf-monitoring/actions/workflows/build.yaml/badge.svg?branch=main)](https://github.com/redhat-appstudio-qe/perf-monitoring/actions/workflows/build.yaml)

This is a monitoring setup for when we run Load/Performance Tests on AppStudio

Runs Prometheus, Pushgateway Grafana and Ingester which communicates with each other to gather and display metrics which are captured during StoneSoup Performance/Load Tests 

Depends upon [concurency-controller](https://github.com/redhat-appstudio-qe/concurency-controller)




