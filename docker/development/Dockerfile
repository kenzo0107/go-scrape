FROM golang:1.12.5-alpine3.9

WORKDIR /go/src/app

ENV GO111MODULE=on

RUN apk add --no-cache \
        alpine-sdk \
        bash \
        tzdata \
        git \
        chromium \
        chromium-chromedriver \
    && go get -u github.com/pilu/fresh

RUN cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

EXPOSE 3000

CMD ["fresh"]
