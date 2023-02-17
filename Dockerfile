FROM golang:1.19.4-alpine3.17
ENV GO111MODULE=on

RUN cd src
RUN mkdir -p github.com/petersephrin/wdbstudio

WORKDIR /go/src/github.com/petersephrin/wdbstudio

COPY ./ ./
RUN go build
CMD ["./wdbstudio"]
