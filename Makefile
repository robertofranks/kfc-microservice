.PHONY: build

api: build
	sam local start-api

build:
	sam build
