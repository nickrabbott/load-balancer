#! /usr/bin/env bash
#
docker-compose stop loadbalancer
docker-compose rm -f loadbalancer
docker images -f "dangling=true"
docker-compose build loadbalancer
docker image prune -f
docker-compose up -d

