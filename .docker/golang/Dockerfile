FROM golang:1.9

ADD . /go/src/github.com/Sharykhin/gl-mail-api

WORKDIR /go/src/github.com/Sharykhin/gl-mail-api

RUN go get .

RUN go install github.com/Sharykhin/gl-mail-api

#ENTRYPOINT /go/bin/gl-mail-api

CMD tail -f /dev/null

EXPOSE 8002