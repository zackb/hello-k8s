.PHONY: build

APP := app

default: build

setup:
	go get -u github.com/go-redis/redis

build: setup
	go build -o bin/$(APP) 

docker:
	docker build -f Dockerfile -t hello-app .

clean:
	rm bin/$(APP)
