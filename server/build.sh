#!/bin/bash

go get -u
go mod tidy

go build -o ../run-server
