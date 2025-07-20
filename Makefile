build:
	go build -o app

runBuild: build
	./app

PHONY: build runBuild
