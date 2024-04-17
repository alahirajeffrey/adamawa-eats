FROM golang:1.22

WORKDIR /adamawa-eats

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

 