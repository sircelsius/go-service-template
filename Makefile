build:
	go build -o /dev/null cmd/app/app.go

test:
	go test ./internal... ./cmd...

lint:
	go run golang.org/x/lint/goling internal... cmd...

vet:
	go vet internal... cmd...

cyclo:
	gocyclo internal/*
