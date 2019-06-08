.PHONY: build

APP := app

default: build

setup:
	go get -u github.com/go-redis/redis

build: setup
	go build -o bin/$(APP) 

kube-deploy: docker
	kubectl apply -f kubernetes.yml

helm-install: docker
	helm install helm/hello-app --namespace hello-ns --name hello

helm-upgrade: docker
	helm upgrade hello helm/hello-app --namespace hello-ns

docker:
	@eval $$(minikube docker-env) ;\
	docker build -f Dockerfile -t hello-app .

clean:
	rm bin/$(APP)
