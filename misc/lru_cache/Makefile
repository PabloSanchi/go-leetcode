run:
	@echo Running the app...
	$(MAKE) build
	@go run cmd/main.go

build:
	@echo Building the app...
	@go build -o bin/app main.go

test:
	@echo Running tests...
	@go test -v ./...

lint:
	@echo Running linter...
	@gofmt -w .
	@echo Done

clean:
	@echo Cleaning build artifacts...
	@rm -rf bin/app
