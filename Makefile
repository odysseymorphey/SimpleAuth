.PHONY: build
build:
	go build -o build/server cmd/main.go

.PHONY: run
run:
	./build/server

.PHONY: clean
clean:
	rm build/server

.DEFAULT_GOAL = build