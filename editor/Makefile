test:
	go test ./... -cover -coverprofile=coverage.txt

coverage: test
	go tool cover -html=coverage.txt

build:
	go build -o editor

lint:
	golint ./...

clean:
	rm -rf main coverage.txt
