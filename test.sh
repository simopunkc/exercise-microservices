#!/bin/bash
WORKDIR=$(pwd)

cd $WORKDIR/exercise-service
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

cd $WORKDIR/user-service
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
