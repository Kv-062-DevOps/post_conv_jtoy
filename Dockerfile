FROM golang:1.14 AS builder

WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go build main.go

FROM ubuntu

ENV POSTPORT=8082
ENV BACKPORT=8083
ENV BACKADDR="127.0.0.1"

EXPOSE ${POSTPORT}

WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]
