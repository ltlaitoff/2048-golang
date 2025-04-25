run-server:
	go run .

# Default run
run:
	go run .

fmt:
	go fmt github.com/ltlaitoff/...

test:
	go test github.com/ltlaitoff/...

build-arch:
	go build .

build-windows:
	env GOOS=windows GOARCH=amd64 go build package-import-path
