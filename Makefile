.PHONY: build deploy

build:
	sam build

deploy:
	sam build
	sam deploy
