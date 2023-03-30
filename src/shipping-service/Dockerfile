FROM golang:1.20.2-alpine3.17 as builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o /go/bin/shipping .

FROM alpine:3.17.2 as release

WORKDIR /src
COPY --from=builder /go/bin/shipping /src/server
ENV APP_PORT=3080

EXPOSE 3080
ENTRYPOINT [ "/src/server" ]

