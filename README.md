# README

This repo creates a custom OpenTelemetry Collector with Stackdriver Exporter.

It followed the tutorial [here](https://medium.com/opentelemetry/building-your-own-opentelemetry-collector-distribution-42337e994b63).

To build and run this project:

1. Ensure you have a GCP project, create a service account, and grant it sufficient permissions to write to [Cloud monitoring](https://cloud.google.com/monitoring).

2. Download a service account key and set your environment variable:

```
export GOOGLE_SERVICE_CREDENTIALS=<path-to-your-sa>
```

3. Change the `exporters.stackdriver.project` field to your GCP project id in config.yaml

4. Start node-exporter locally so we can have some metrics avaiable:

```
docker run -p 9100:9100 --name node-exporter-${RANDOM} bitnami/node-exporter:latest
# you can test the metrics endpoint as well
curl -s http://localhost:9100/metrics | head
```

5. Start the collector

```
go run . --config config.yaml
```

6. Go to the [Metrics explorer](https://pantheon.corp.google.com/monitoring/metrics-explorer) and search for some metrics from node-Exporter, e.g. `node_cpu_seconds_total`.
