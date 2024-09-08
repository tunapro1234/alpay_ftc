# Define the default target
.PHONY: all
all: attach

# Build and run the development container in detached mode
.PHONY: build
build:
	docker compose up --build -d

# Attach to the running development container
.PHONY: attach
attach:
	docker compose start
	docker exec -it alpay_dev /bin/bash

# Attach to the running development container
.PHONY: stop
stop:
	docker stop -t 0 alpay_dev

# Build the release Docker image
.PHONY: release-build
release-build:
	docker build -t alpay_release -f docker/release_container/Dockerfile .

