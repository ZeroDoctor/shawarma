
.PHONY: generate
generate:
	@echo generate...
	pkl-gen-go shawarma.pkl

.PHONY: build
build:
	@echo building...
	go build -o shawarma ./cmd/main.go

.PHONY: run
run:
	@echo running...
	./shawarma
