clean:
	rm -rf bin/*

build-binary:
	go build -o bin/shorten-url .

build-docker:
	docker build -t shorten-url-go:latest .

run:
	go run main.go