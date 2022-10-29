#!/bin/bash
docker compose down
docker compose rm -f
sleep 5
docker volume rm $(docker volume ls -q)
sleep 5
