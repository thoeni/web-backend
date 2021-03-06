FROM golang:1.5

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

COPY ./src /go/src/app
RUN go-wrapper download
RUN go-wrapper install

EXPOSE  8080

CMD ["go-wrapper", "run"]

