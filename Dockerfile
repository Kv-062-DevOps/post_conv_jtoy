FROM golang:1.14 AS builder

WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go build 

FROM ubuntu
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]
