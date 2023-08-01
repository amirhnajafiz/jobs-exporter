# Stallion Vet

Monitoring **Stallion** message broker server with a Blackbox exporter. 
Check service health with metrics like response time, 
success publish, failure requests, and failure connections.

## Start

Stallion server set up:

```shell
go run server/main.go
```

Start Blackbox exporter:

```shell
go run main.go
```

Check metrics:

```shell
curl -i localhost:3030/metrics
```

## Configs

```yaml
client:
  address: "127.0.0.1"
  port: 9090
  auth:
    username: "root"
    password: "Pa$$word"
telemetry:
  address: ":3030"
  enabled: true
```
