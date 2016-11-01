FROM alpine
MAINTAINER Calvin Leung Huang <https://github.com/calvn>

COPY build/proto /

RUN apk update
RUN apk --no-cache --no-progress add ca-certificates \
    && rm -rf /var/cache/apk/*

ENTRYPOINT ["/proto"]
