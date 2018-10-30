FROM golang:1.11-alpine as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on

WORKDIR $GOPATH/src/github.com/castaneai/go-spanner

RUN apk update && apk add git

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go install -ldflags="-w -s"

FROM scratch
COPY --from=builder /go/bin/go-spanner /go/bin/go-spanner
ENTRYPOINT ["/go/bin/go-spanner"]