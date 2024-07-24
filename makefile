COMPOSE := docker-compose -f docker-compose.yml


# Infrastructure
build:
	docker build -t $(IMAGE):$(TAG) .

compose-up:
	@echo "🚢 Starting services..."
	@$(COMPOSE) down -v || true
	@$(COMPOSE) up -d --force-recreate

compose-down:
	@$(COMPOSE) down -v || true
	