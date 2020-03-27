FROM golang:1.14 AS builder

EXPOSE 8082

ENV GO111MODULE=on
ARG PORT=":8082"
ENV PORT="${PORT}"
ARG DBLINK="http://127.0.0.1:8083/add"
ENV DBLINK="${DBLINK}"

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
