FROM golang:alpine AS build-env
ADD . /go/src/github.com/mdanzinger/edealer-cms/admin/
RUN apk add --no-cache git mercurial \
#    && go get -u github.com/golang/dep/... \
    && go get github.com/gorilla/mux \
#    && go get github.com/jinzhu/gorm \
#    && go get github.com/go-sql-driver/mysql \
#    && go get github.com/gomodule/redigo/redis \
#    && go get github.com/satori/go.uuid \
#    && go get golang.org/x/crypto/bcrypt \
#    && cd /go/src/github.com/mdanzinger/ed-status/backend/ && dep ensure \
    && apk del git mercurial
RUN cd /go/src/github.com/mdanzinger/edealer-cms/admin/ && go build -o goapp


FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/github.com/mdanzinger/edealer-cms/admin/goapp /app/
#COPY --from=build-env /go/src/github.com/mdanzinger/edealer-cms/admin/templates /app/templates
RUN ["chmod", "+x", "./goapp"]
ENV PORT 8090
ENTRYPOINT ./goapp
