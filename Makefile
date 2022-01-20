.PHONY = build test

GRAFANA_DIR=./grafana
LOG_DIR=./logs
CHECK_DIR=go list ./... | grep -v /cmd/utilits
SQL_DIR=./scripts
MICROSERVICE_DIR=$(PWD)/internal/microservices

stop-redis:
	systemctl stop redis
stop-postgres:
	systemctl stop postgresql
run-posgres-redis:
	systemctl start redis
	systemctl start postgresql

build:
	go build -o server.out -v ./cmd/server

build-docker:
	docker build -t thecompiler .

run:
	docker run  --memory 2G --log-opt max-size=5M --log-opt max-file=3 -p 5000:5000 -p 5432:5432 --name thecompiler -t thecompiler

run-build: build-docker run

open-last-log:
	cat $(LOG_DIR)/`ls -t $(LOG_DIR) | head -1 `

clear-logs:
	rm -rf $(LOG_DIR)/*.log

rm-docker:
	docker rm -vf $$(docker ps -a -q) || true

build-utils:
	go build -o utils.out -v ./cmd/utilits

parse-last-log: build-utils
	./utils.out -search-url=${search_url}

gen-mock:
	go generate -n $$(go list ./internal/...)

test:
	go test -v -race ./internal/...

