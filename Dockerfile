FROM golang:1.12-alpine as build
ENV GO111MODULE=on
WORKDIR $GOPATH/src/github.com/joecomscience/alert-manager
COPY . .
RUN apk update && \
    apk add git --no-cache && \
    go get ./... && go build main.go

FROM alpine:3 as release

COPY --from=build /go/src/github.com/joecomscience/alert-manager/main .
CMD [ "./main" ]