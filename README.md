# Hello Kubernetes
### Bare bones Kubernetes app deployment
- Go application that serves HTTP via `PORT` env var
  - built and run with single Dockerfile using the Makefile
- Redis server
- Minikube

### Run
1. Start minikube
   1. `minikube start`
   2. use minikube's docker env `eval $(minikube docker-env)`
   2. switch to minikube's context `kubectx minikube`
2. Build image
   1. `make docker`
      1. adds the image to the minikube docker registry
3. Deploy kubernetes app
   1. With kubectl: `kubectl apply -f kubernetes.yml`
   2. OR with helm: `make helm-install` or `make helm-upgrade`
4. Update hosts file (ingress uses hostname: `hello.app`)
   1. CAREFUL: `echo $(minikube ip) hello.app >> /etc/hosts` 
5. Open app: http://hello.app
