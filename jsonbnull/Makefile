default: test

test:
	go test ./...


startdb:
	docker run -d --rm \
			--name=demoDB \
			-p 5432:5432 \
			-e POSTGRES_DB=jsonbnull \
			-e POSTGRES_USER=jsonbnull \
			-e POSTGRES_PASSWORD=jsonbnull postgres:12-alpine

stopdb:
	docker stop demoDB

.SILENT: startdb stopdb