FROM golang:1.12-alpine as builder

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

FROM alpine
RUN apk add --update --no-cache ca-certificates
COPY --from=builder /go/bin/go-spanner /bin/go-spanner
ENTRYPOINT ["/bin/go-spanner"]