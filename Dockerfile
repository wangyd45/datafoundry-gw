FROM registry.new.dataos.io/lixw/golang:latest
EXPOSE 10012
ADD . /go/src/github.com/asiainfoLDP/datafoundry-gw
WORKDIR /go/src/github.com/asiainfoLDP/datafoundry-gw
RUN go build -o server main.go
RUN chmod +x server
EXPOSE 10012
CMD ["./server"]
