PWD := $(shell pwd)

default: test

start-db:
	docker run -d --rm \
			--name=demo-db \
			-p 5444:5432 \
			-e POSTGRES_DB=demo_test \
			-v $(PWD)/db:/docker-entrypoint-initdb.d \
			-e POSTGRES_PASSWORD=postgres postgres:12-alpine

stop-db:
	docker stop demo-db || echo ""

test: start-db
	go test .

.SILENT: stop-db