
.PHONY: build
build:
	@echo building...
	pkl-gen-go template.pkl
	go build -o shawarma ./cmd/main.go

.PHONY: run
run:
	./shawarma
