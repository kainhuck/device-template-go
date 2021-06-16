ARG BASE=golang:1.15-alpine3.12
FROM ${BASE} AS builder

ARG MAKE='make cmd/device-template'

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories

RUN apk add --update --no-cache make git

# set the working directory
WORKDIR /device-template-go

COPY . .
RUN --mount=type=cache,target=/go,id=coap_driver_cache,sharing=shared go mod download
RUN --mount=type=cache,target=/root/.cache,id=coap_driver_build_cache,sharing=shared ${MAKE}

FROM alpine:3.12

EXPOSE 61618

COPY --from=builder /device-template-go/cmd/device-template /bin/
COPY --from=builder /device-template-go/cmd/res /etc/driver/res/

ENTRYPOINT ["/bin/device-template"]
CMD ["--confdir=/etc/driver/res"]
