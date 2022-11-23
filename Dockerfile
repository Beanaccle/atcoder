FROM golang:1.14.1-alpine
RUN apk --update-cache --no-cache add \
  git \
  nodejs \
  npm \
  py3-pip && \
  pip3 install --upgrade pip && \
  pip3 install online-judge-tools && \
  npm install -g atcoder-cli

WORKDIR /go/src
COPY /templates/golang /root/.config/atcoder-cli-nodejs/golang
RUN acc config default-template golang

ENV GO111MODULE on
RUN go get github.com/nsf/gocode \
  github.com/uudashr/gopkgs/v2 \
  github.com/go-delve/delve/cmd/dlv \
  golang.org/x/tools/gopls@v0.9.5
