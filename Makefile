GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")

cli:
	go build -mod $(GOMOD) -ldflags="-s -w" -o bin/copy-uri cmd/copy-uri/main.go
	go build -mod $(GOMOD) -ldflags="-s -w" -o bin/copy cmd/copy/main.go
