FROM golang:1.14.2-alpine3.11
RUN apk add --no-cache openssh-client
COPY id_rsa /root/.ssh/
COPY known_hosts /root/.ssh/
COPY gitconfig /root/.gitconfig
COPY config /root/.ssh/


RUN apk update && apk --no-cache add ca-certificates

RUN apk update && apk add tzdata
RUN cp /usr/share/zoneinfo/Asia/Kolkata /etc/localtime
RUN echo "Asia/Kolkata" > /etc/timezone

RUN apk update && apk add -f git librdkafka-dev pkgconf build-base
RUN mkdir -p /platform/campaign_service/
ADD . /platform/campaign_service/

ENV SERVICE=campaign_service
ENV NAMESPACE=/platform
ENV CONFIG_DIR=/platform/campaign_service/core/golang/config
ENV ENV=dev
WORKDIR /platform/campaign_service/zerotouch/golang

RUN go build -o main .

EXPOSE 7800
EXPOSE 7805
CMD ["/platform/campaign_service/zerotouch/golang/main"]
