FROM golang:1.13 as builder

RUN git config --global core.eol lf
RUN git config --global core.autocrlf input

COPY . /go/src/github.com/kucinghitam13/alien
WORKDIR /go/src/github.com/kucinghitam13/alien

# building binary file
RUN go get -v -u github.com/golang/dep/cmd/dep
RUN dep ensure -v
RUN go build -v -o ./build/alien cmd/*/*.go
FROM alpine:3.6 as alpine

RUN apk add -U --no-cache ca-certificates

FROM ubuntu

COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /go/src/github.com/kucinghitam13/alien/build/alien ./build/
WORKDIR /build

EXPOSE 1981

ENTRYPOINT ["./alien"]
