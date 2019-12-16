build:
	go build -o /dev/null cmd/app/app.go

lint:
	golint internal/... cmd/...

vet:
	go vet internal/... cmd/...

cyclo:
	gocyclo internal/*
