.PHONY: run/hello
run/hello:
	@go run ./Hello/main.go
	
.PHONY: build/hello
build/hello:
	@go build -o ./bin/hello ./Hello/main.go

.PHONY: run/stringers
run/stringers:
	@go run ./Stringers/main.go
	
.PHONY: build/stringers
build/stringers:
	@go build -o ./bin/stringers ./Stringers/main.go

.PHONY: run/fibonacci
run/fibonacci:
	@go run ./Fibonacci/main.go
	
.PHONY: build/fibonacci
build/fibonacci:
	@go build -o ./bin/fibonacci ./Fibonacci/main.go

.PHONY: run/sqrt
run/sqrt:
	@go run ./Sqrt/main.go
	
.PHONY: build/sqrt
build/sqrt:
	@go build -o ./bin/sqrt ./Sqrt/main.go