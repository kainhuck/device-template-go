ARG BASE=golang:1.16-alpine3.12
FROM ${BASE} AS builder

RUN sed -e 's/dl-cdn[.]alpinelinux.org/nl.alpinelinux.org/g' -i~ /etc/apk/repositories
RUN apk add --update --no-cache make git openssh gcc libc-dev zeromq-dev libsodium-dev

# set the working directory
WORKDIR /device-template-go

COPY . .

RUN go mod tidy
RUN go mod download

RUN make build

FROM alpine:3.12

RUN sed -e 's/dl-cdn[.]alpinelinux.org/nl.alpinelinux.org/g' -i~ /etc/apk/repositories
RUN apk add --update --no-cache zeromq dumb-init

COPY --from=builder /device-template-go/cmd /

EXPOSE 59982

ENTRYPOINT ["/device-template"]
CMD ["--cp=consul://edgex-core-consul:8500", "--registry", "--confdir=/res"]
