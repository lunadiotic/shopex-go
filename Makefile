# Define the default target that runs when someone types "make" without an arguments.
.PHONY: help
help:
	@echo "Makefile for E-commerce Backend."
	@echo ""
	@echo "Type 'make help' to show this help message."
	@echo "Type 'make run' to run the application."

# The 'run' command compiles and starts our API entry point.
.PHONY: run
run:
	# Run the application
	go run cmd/api/main.go