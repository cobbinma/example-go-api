FROM golang:1.14-stretch

ENV GO111MODULE=on
WORKDIR /go/src

COPY go.mod .
COPY go.sum .
RUN go mod download

RUN ["go", "get", "github.com/cespare/reflex"]

COPY . .

ENTRYPOINT ["reflex", "-c", "./reflex.conf"]
