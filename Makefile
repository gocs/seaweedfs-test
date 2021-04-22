
.PHONY: run
run:
	@GOOS="linux" GOARCH="amd64" CGO_ENABLED=0 go build -o app main.go
	@docker-compose down && docker-compose up --build -d