FROM golang:alpine AS builder
RUN apk --no-cache add git curl make glide
WORKDIR /go/src/github.com/rschmicker/honeypot
COPY . .
RUN make

FROM alpine:latest
WORKDIR /
COPY --from=builder /go/bin/honeypot /bin/
EXPOSE 2222
ENTRYPOINT ["honeypot"]