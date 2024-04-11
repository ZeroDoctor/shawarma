
.PHONY: generate
generate:
	@echo generate...
	cd server && \
	pkl-gen-go ../shawarma.pkl

.PHONY: build
build:
	@echo building...
	cd server && \
	go build -o ../shawarma ./cmd/main.go

.PHONY: run
run:
	@echo running...
	./shawarma
