FROM golang:alpine

RUN mkdir /app
WORKDIR /app

COPY . .

RUN go mod download

CMD ["go", "run", "cmd/gowait.go"]