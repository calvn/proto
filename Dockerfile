FROM alpine
MAINTAINER Calvin Leung Huang <https://github.com/calvn>

COPY build/proto /

ENTRYPOINT ["/proto"]
