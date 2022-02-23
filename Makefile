.PHONY: run/hello
run/hello:
	@go run ./Hello/main.go
	
.PHONY: build/hello
build/hello:
	@go build -o ./bin/hello ./Hello/main.go

.PHONY: run/stringers
run/hello:
	@go run ./Stringers/main.go
	
.PHONY: build/stringers
build/hello:
	@go build -o ./bin/stringers ./Stringers/main.go
