clean:
	rm -rf bin/* ;\
	rm -rf coverage/* ;\

build-binary:
	go build -o bin/shorten-url .

build-docker:
	docker build -t shorten-url-go:latest .

ci-test:
	go test -v ./... -covermode=count -coverprofile=coverage/coverage.out ;\
	go tool cover -func=coverage/coverage.out -o=coverage/coverage.out ;\

test:
	go test ./... -coverprofile=coverage/coverage.out ;\
	go tool cover -func=coverage/coverage.out ;\
	go tool cover -html=coverage/coverage.out ;\

run:
	go run main.go