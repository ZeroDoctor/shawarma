
.PHONY: init
init:
	go install github.com/swaggo/swag/cmd/swag@latest

.PHONY: shawarma-generate
	@echo generate pkl...
	cd server && \
	pkl-gen-go ../shawarma.pkl

.PHONY: swag-generate
swag-generate:
	@echo generate swag...
	cd ./server && \
	swag fmt && \
	swag init -g ./internal/api/api.go --parseInternal --pdl 2 --pd 

.PHONY: generate
generate: shawarma-generate

.PHONY: build
build: swag-generate
	@echo building server...
	cd server && \
	go build -o ../shawarma ./cmd/main.go

.PHONY: run
run:
	@echo running server...
	./shawarma

.PHONY: clean
clean:
	@echo cleaning server...
	rm shawarma || true
	rm shawarma.db || true

.PHONY: ui 
ui:
	@echo running ui...
	cd ui && \
	npm run dev
