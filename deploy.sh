
oc create ns performance-monitoring --dry-run=client  -o yaml | oc apply -f - 


oc create configmap grafana-dashboard-config \
    -n performance-monitoring \
    --from-file=monitoring/grafana/provisioning/dashboards/dashboard.yml \
    --from-file=monitoring/grafana/provisioning/dashboards/loadtest.json \
    --dry-run=client -o yaml | oc apply -f -

oc create configmap grafana-dashboard-datasources \
    -n performance-monitoring \
    --from-file=monitoring/grafana/provisioning/datasources/ds.yml \
    --dry-run=client -o yaml | oc apply -f -

oc apply -f openshift/