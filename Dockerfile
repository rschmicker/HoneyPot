FROM golang:1.11-alpine
RUN apk --no-cache add git make curl glide
WORKDIR /go/src/github.com/rschmicker/HoneyPot
COPY . .
RUN make
