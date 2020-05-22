FROM golang AS builder
MAINTAINER lucmichalski <michalski.luc@gmail.com>

# RUN apk add --no-cache make gcc g++ ca-certificates musl-dev make git

COPY . /go/src/github.com/x0rzkov/gocolly-kafka-docker/cmd/producer
WORKDIR /go/src/github.com/x0rzkov/gocolly-kafka-docker/cmd/producer

RUN go install

FROM debian:stable-slim AS runtime
MAINTAINER lucmichalski <michalski.luc@gmail.com>

WORKDIR /opt/lucmichalski/bin

# copy executable
COPY --from=builder /go/bin/producer /opt/lucmichalski/bin/gocolly-producer

ENV PATH $PATH:/opt/lucmichalski/bin

# Container configuration
# EXPOSE 9000
CMD ["/opt/lucmichalski/bin/gocolly-producer"]


