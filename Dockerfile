FROM golang:1.9.2

EXPOSE 10012

COPY . /go/src/github.com/asiainfoLDP/datafoundry-gw

WORKDIR /go/src/github.com/asiainfoLDP/datafoundry-gw

# NOT DO RUN go run main.go
RUN go build

CMD ./datafoundry-gw