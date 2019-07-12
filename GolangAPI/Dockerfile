FROM golang:1.12.0-alpine


WORKDIR /go/src/github.com/s14228so/WilicoExtra/GolangAPI


RUN cd /go/src/github.com/s14228so/WilicoExtra/GolangAPI \
    && ls -la \ 
    && apk update \
    && apk add --no-cache git \
    && go get -u github.com/codegangsta/gin \
    && go get -u github.com/golang/dep/cmd/dep \
    && dep ensure 
