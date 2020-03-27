FROM golang:1.14 AS builder

WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go build 

FROM ubuntu

ENV PORT=":8082"
ENV DBLINK="http://127.0.0.1:8083/add"

EXPOSE 8082

WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]
