
.PHONY: build
build:
	@echo building...
	pkl-gen-go example/template.pkl
	go build -o shawarma ./cmd/main.go
