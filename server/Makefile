docs:
	go install github.com/swaggo/swag/cmd/swag@latest && \
	export PATH=$(go env GOPATH)/bin:$PATH && \
	swag init -g cmd/api/main.go

build:
	CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/api/main.go

start_dev_db:
	docker-compose down && \
	docker-compose up

start_dev_server: 
	export GO_ENV=development && go run ./cmd/api/main.go

start_dev: start_dev_db start_dev_server