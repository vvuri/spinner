.PHONY: help
help:
	@echo "API server on port 80"

run:
	go run ./cmd/main/app.go

docker-build:
	docker build -t gometrics:v1 .

docker-run:
	docker run --rm -p 8080:80 gometrics:v1

docker-run-sh:
	docker run -it --rm gometrics:v1 sh
