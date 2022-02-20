.PHONY: run/hello
run/hello:
	@go run ./Hello/main.go
	
.PHONY: build/hello
build/hello:
	@go build -o ./bin/hello ./Hello/main.go