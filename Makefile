APP_NAME := Calculator with REST API
ENTRYPOINT := cmd/main.go

.PHONY: docs

docs:
	swag init -g $(ENTRYPOINT)