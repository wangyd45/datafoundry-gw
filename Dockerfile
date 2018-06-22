FROM registry.new.dataos.io/lixw/golang:latest

ADD . /go/src/github.com/asiainfoLDP/datafoundry-gw
WORKDIR /go/src/github.com/asiainfoLDP/datafoundry-gw
RUN go build -o server main.go && chmod +x server
EXPOSE 10012
ENTRYPOINT ["./server"]
