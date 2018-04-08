FROM golang:1.9.2

EXPOSE 10012

COPY . /src/github.com/asiainfoLDP/datafoundry-gw

WORKDIR /src/github.com/asiainfoLDP/datafoundry-gw

RUN go run main.go