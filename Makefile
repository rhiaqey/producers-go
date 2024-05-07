export REDIS_VERSION=7.2.4
export REDIS_PASSWORD=welcome

.PHONY: run
run:
	ID=1 \
	NAME=ticker \
	NAMESPACE=rhiaqey \
		go run bin/ticker.go

.PHONY: dev
dev: build

.PHONY: build
build:
	@go build -o out/ticker bin/ticker.go
	ls -lah out/ticker

.PHONY: prod
prod: build
	@go build -o out/ticker -ldflags "-s -w" bin/ticker.go
	ls -lah out/ticker

.PHONY: hub
hub:
	docker run -it --rm --name hub -p 6379:6379 \
		-e REDIS_PASSWORD=${REDIS_PASSWORD} \
		-e PUBLIC_KEY=/certs/pub.pem \
		-e PRIVATE_KEY=/certs/priv.pem \
		-e PRIVATE_PORT=3000 \
		-e PUBLIC_PORT=3001 \
		-v ./certs:/certs \
		--entrypoint 'rhiaqey-hub' \
		--network host \
		rhiaqey/hub:latest 'run'

.PHONY: redis
redis:
	docker run -it --rm --name redis -p 6379:6379 \
		-e ALLOW_EMPTY_PASSWORD=no \
		-e REDIS_PASSWORD=${REDIS_PASSWORD} \
		--network host \
		rhiaqey/redis:${REDIS_VERSION}
