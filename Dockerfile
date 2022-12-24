FROM alpine:latest as kotlin

ENV KOTLIN_VERSION 1.3.71
RUN apk --no-cache add wget && \
  cd /usr/local && \
  wget -q https://github.com/JetBrains/kotlin/releases/download/v${KOTLIN_VERSION}/kotlin-compiler-${KOTLIN_VERSION}.zip && \
  unzip kotlin-compiler-*.zip && \
  rm kotlin-compiler-*.zip && \
  rm -f kotlinc/bin/*.bat

FROM ruby:2.7.1-alpine

COPY --from=golang:1.14.1-alpine /usr/local/go/ /usr/local/go/
ENV PATH /usr/local/go/bin:$PATH

COPY --from=kotlin /usr/local/kotlinc /usr/local/kotlinc
ENV PATH /usr/local/kotlinc/bin:$PATH

RUN apk --no-cache add \
  openjdk11 --repository http://dl-cdn.alpinelinux.org/alpine/edge/community \
  bash \
  nodejs \
  npm \
  py3-pip && \
  pip3 install --upgrade pip && \
  pip3 install online-judge-tools && \
  npm install -g atcoder-cli && \
  acc config default-test-dirname-format test

WORKDIR /usr/src/app
COPY /templates /root/.config/atcoder-cli-nodejs
# RUN acc config default-template golang

ENV GO111MODULE on
RUN go get github.com/nsf/gocode \
  github.com/uudashr/gopkgs/v2 \
  github.com/go-delve/delve/cmd/dlv \
  golang.org/x/tools/gopls@v0.9.5 && \
  apk --no-cache add \
  make \
  gcc \
  musl-dev && \
  gem install solargraph

CMD ["/bin/bash"]
