
.PHONY: generate
generate:
	@echo generate...
	pkl-gen-go template.pkl

.PHONY: build
build:
	@echo building...
	go build -o shawarma ./cmd/main.go

.PHONY: run
run:
	@echo running...
	./shawarma
