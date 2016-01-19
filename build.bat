	cd cmd/server; \
	go build
	cd ../..
	cd service; \
	go build
	cd ..
	cd api; \
	go build