#!/usr/bin/env bash

# This file prevents go from caching test results and force a run every time a test needs it.
## Runs project coverage
## Detects race conditions
go test ./... --count=1 -coverprofile coverage.out -coverpkg=./... -race

go tool cover -html=coverage.out