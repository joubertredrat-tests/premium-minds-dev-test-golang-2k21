coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out

build-mac:
	env GOOS=darwin GOARCH=amd64 go build -o pokemon-darwin-amd64

build-linux:
	go build -o pokemon-linux-amd64