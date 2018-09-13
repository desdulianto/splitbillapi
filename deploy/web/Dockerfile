FROM golang:1.10-alpine3.8 AS compile

COPY . /go/src/github.com/desdulianto/splitbillapi

RUN apk --no-cache add curl make \
    && curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN cd /go/src/github.com/desdulianto/splitbillapi \
    && make dep \
    && make compile 

FROM alpine:3.8

COPY --from=compile /go/src/github.com/desdulianto/splitbillapi/deploy/_output/web/bin/web /usr/local/bin/splitbillapi

EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/splitbillapi"]
