FROM alpine:latest

MAINTAINER thomas91310

WORKDIR "/opt"

ADD .docker_build/bot-playground /opt/bin/bot-playground

CMD ["/opt/bin/bot-playground"]
