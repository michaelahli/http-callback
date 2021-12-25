FROM golang:1.17.0-alpine3.13

RUN apk update && apk add bash
RUN set -ex && apk --no-cache add sudo

WORKDIR /app/http-callback

COPY . .
COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go build -o binary

CMD ["/app/http-callback/binary"] 