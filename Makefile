run: build
	@./bin/pricefetcher

build:
	@go build -o bin/pricefetcher . 
