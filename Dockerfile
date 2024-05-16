FROM golang:1.22-bullseye AS builder

WORKDIR /

COPY go.mod go.sum ./

RUN go mod download && go mod verify

ADD . .

RUN go build -o /todo-app

EXPOSE 8080

CMD ["/todo-app"]