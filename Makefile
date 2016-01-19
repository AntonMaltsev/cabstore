build:
	cd cmd/server; \
	go build

deps:
	cd cmd/server; \
	go get

test:
	./cmd/server/server --config config.yaml server & \
	pid=$$!; \
	go test; \
	kill $$pid