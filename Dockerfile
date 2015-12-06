FROM golang:1.5

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

COPY ./src /go/src/app
RUN go-wrapper download
RUN go-wrapper install

EXPOSE  8080

# this will ideally be built by the ONBUILD below ;)
CMD ["go-wrapper", "run"]

