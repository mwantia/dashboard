# Makefile for Admin Dashboard
BINARY_NAME := dashboard

.PHONY: all build run clean serve help

all: clean run

build:
	@echo "Building go..."
	go build -o ./build/$(BINARY_NAME) ./cmd/server/main.go

run: build
	@echo "Running server..."
	./build/$(BINARY_NAME)

clean:
	go clean
	rm -f ./build/$(BINARY_NAME)

serve:
	@air -c .air.toml

help:
	@echo "Available commands:"
	@echo "  make build   - Build the application"
	@echo "  make run     - Run the application"
	@echo "  make clean   - Remove binaries and application clean up"
	@echo "  make serve   - Run air on port 9001 for live reloads with templ and tailwindcss"