from golang:stretch

WORKDIR /go/src/github.com/zackb/hello-k8s

COPY . ./

RUN make

CMD bin/app

