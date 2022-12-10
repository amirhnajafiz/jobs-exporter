<h1 align="center">
Stallion Black Box Exporter
</h1>

Monitoring Stallion message broker with a Black Box. 
Check service health with metrics like response time, 
success publish and failure requests.

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
