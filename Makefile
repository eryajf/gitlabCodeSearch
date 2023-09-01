default: build

build:
	go build -o eryajfctl main.go

build-linux:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o eryajfctl main.go

lint:
	env GOGC=25 golangci-lint run --fix -j 8 -v ./...