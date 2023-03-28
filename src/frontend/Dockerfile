FROM golang:1.20.2-alpine3.17 as builder

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o /go/bin/frontend .

FROM alpine:3.17.2 as release
WORKDIR /src
COPY --from=builder /go/bin/frontend /src/server
COPY ./templates ./templates
COPY ./static ./static
COPY .env .env

EXPOSE 8080
ENTRYPOINT [ "/src/server" ]
