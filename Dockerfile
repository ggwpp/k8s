FROM golang:1.10-alpine
LABEL maintainer="Grean-Developers-Family"

RUN apk add --no-cache git mercurial \
    && go get github.com/gorilla/mux \
    && apk del git mercurial

WORKDIR /go/src/app
ADD src/ .
RUN go build 

CMD [ "./app" ]
EXPOSE 8080