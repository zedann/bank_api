build:
	@go build ./cmd/app/main.go

run: build
	@./main

air:
	air --build.cmd "go build -o bin/api cmd/app/main.go" --build.bin "./bin/api"
	