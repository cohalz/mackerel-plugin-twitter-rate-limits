deps:
	dep ensure -vendor-only

build: deps
	GOARCH=amd64 GOOS=linux go build -o bin/rate_limits_checker
