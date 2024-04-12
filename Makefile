GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")

cli:
	go build -mod $(GOMOD) -ldflags="-s -w" -o bin/copy-uri cmd/copy-uri/main.go
	go build -mod $(GOMOD) -ldflags="-s -w" -o bin/copy cmd/copy/main.go

lambda:
	@make lambda-copy-uri

lambda-copy-uri:
	if test -f bootstrap; then rm -f bootstrap; fi
	if test -f copy_uri.zip; then rm -f copy_uri.zip; fi
	GOARCH=arm64 GOOS=linux go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -tags lambda.norpc -o bootstrap cmd/copy-uri/main.go
	zip copy_uri.zip bootstrap
	rm -f bootstrap
