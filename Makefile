APP_NAME := app
BIN_DIR := bin
MAIN_FILE := ./cmd/main.go
ENV_FILE := .env

.DEFAULT_GOAL := help

build:
	@echo "📦 Compilando o projeto..."
	@mkdir -p $(BIN_DIR)
	@go build -o $(BIN_DIR)/$(APP_NAME) $(MAIN_FILE)
	@echo "✅ Binário gerado em $(BIN_DIR)/$(APP_NAME)"

run: build
	@echo "🏃 Executando a aplicação..."
	@$(BIN_DIR)/$(APP_NAME)

docker-build:
	@echo "🐳 Buildando Docker..."
	@docker compose -f infra/docker-compose.yaml up -d --build

up:
	@echo "⬆️  Subindo containers..."
	@docker compose -f infra/docker-compose.yaml up -d --build

clean:
	@echo "🧹 Limpando binário..."
	@rm -rf $(BIN_DIR)

rebuild: clean build

help:
	@echo ""
	@echo "🧰 Comandos disponíveis:"
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z\-\_]+:.*##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } ' $(MAKEFILE_LIST)
