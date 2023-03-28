FROM golang:1.20.2-alpine3.17 as builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o /go/bin/productcatalog .

FROM alpine:3.17.2 as release

WORKDIR /src
COPY --from=builder /go/bin/productcatalog /src/server
COPY products.json .

EXPOSE 3050
ENTRYPOINT [ "/src/server" ]

