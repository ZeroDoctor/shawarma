
.PHONY: generate
generate:
	@echo generate pkl...
	cd server && \
	pkl-gen-go ../shawarma.pkl

.PHONY: build
build:
	@echo building server...
	cd server && \
	go build -o ../shawarma ./cmd/main.go

.PHONY: run
run:
	@echo running server...
	./shawarma

.PHONY: ui 
ui:
	@echo running ui...
	cd ui && \
	npm run dev &
