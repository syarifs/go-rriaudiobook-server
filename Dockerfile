FROM golang:1.18 as build-env

WORKDIR /go/src/app
ADD . /go/src/app

RUN go get -d -v ./...

RUN go build -o /go/bin/app

FROM scratch as prod-env
COPY --from=build-env /go/bin/app /
EXPOSE 8080
CMD ["/app"]
