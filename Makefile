build:
	go build -o ./bin/modularis

run: build
	./bin/modularis

test:
	go test ./...