APP_NAME := Calculator with REST API
ENTRYPOINT := cmd/main.go

.PHONY: docs tests cover

docs:
	swag init -g $(ENTRYPOINT)

tests:
	go test ./... -coverprofile=covarage.out

cover:
	go tool cover -html=covarage.out -o coverage.html
	firefox coverage.html