PKGS := $(shell go list ./... | grep -v /vendor)

daily:
	CGO_ENABLED=0 go build -o bin/daily ./pkg/main/


