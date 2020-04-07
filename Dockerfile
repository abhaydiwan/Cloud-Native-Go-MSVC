FROM golang:1.12.0-alpine3.9
RUN apk update && apk upgrade && apk add --no-cache bash git openssh
RUN go get github.com/gin-gonic/gin

ENV SOURCES /go/src/github.com/abhaydiwan/Cloud-Native-Go-MSVC/
COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go build

WORKDIR ${SOURCES}

CMD ${SOURCES}Cloud-Native-Go-MSVC
EXPOSE 8090
