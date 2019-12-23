build:
	go build -o /dev/null cmd/app/app.go

test:
	go test ./internal... ./cmd...

lint:
	go run golang.org/x/lint/golint ./internal... ./cmd...

vet:
	go vet internal... cmd...

cyclo:
	go run github.com/fzipp/gocyclo internal/*

imports:
	 go run golang.org/x/tools/cmd/goimports -w ./internal
