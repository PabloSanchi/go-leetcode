run:
	@$(MAKE) build
	@echo "Running the app..."
	@./bin/app

build:
	@echo "Building the app..."
	@touch splitwise.db
	@go build -o bin/app cmd/main.go

test:
	@echo "Running tests..."
	@go test -v ./...

lint:
	@echo "Running linter..."
	@gofmt -w .

clean:
	@echo "Cleaning build artifacts..."
	@rm -rf bin
	@rm -rf splitwise.db
