FROM golang:1.20.2-alpine3.17 as builder

RUN apk add --no-cache git

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build publisher.go

FROM alpine:3.17.2 as release

WORKDIR /app
COPY --from=builder /src/publisher /app/publisher

CMD [ "/app/publisher" ]

