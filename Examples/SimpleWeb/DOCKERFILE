FROM golang

ENV GO111MODULE=on
ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct

RUN mkdir -p /root/gopath
RUN mkdir -p /root/gopath/src
RUN mkdir -p /root/gopath/pkg
RUN mkdir -p /root/gopath/bin
RUN mkdir -p /root/code
ENV GOPATH /root/gopath


WORKDIR /godocker

ADD . /godocker

RUN go build -v -o app main.go

EXPOSE 8080

ENTRYPOINT ["./app"]