# Simple load balancer project in Go

I wrote this to get some reps with Go and Docker. This implements a simple round robin balancing algorithm for two web servers.

## Installtion

Ensure Docker and Go are installed.

```bash
docker-compose up -d
```

### Rebuild

When writing this I used a script that would only build the Go container and wouldn't rebuild the apache containers.

```bash
#! /usr/bin/env bash
#
docker-compose stop loadbalancer
docker-compose rm -f loadbalancer
docker images -f "dangling=true"
docker-compose build loadbalancer
docker image prune -f
docker-compose up -d
```
