FROM golang:1.12-alpine as build
ENV GO111MODULE=on
WORKDIR $GOPATH/src/github.com/joecomscience/alert-manager
COPY . .
RUN apk update && \
    apk add git --no-cache && \
    apk --no-cache add ca-certificates && \
    go get ./... && \
    go build main.go

FROM alpine:3 as release

ENV PORT=8080

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /go/src/github.com/joecomscience/alert-manager/main .
CMD [ "./main" ]