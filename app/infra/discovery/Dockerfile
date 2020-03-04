FROM golang:latest

MAINTAINER  CloudZou
WORKDIR /go/src/
RUN mkdir /go/src/app/
RUN mkdir /go/src/script/
ADD ./cmd/discovery/main /go/src/app/
ADD /script/build.sh /go/src/script/
ADD /script/discovery.toml /go/src/app/discovery.toml
EXPOSE 7171
CMD ["/bin/bash", "/go/src/script/build.sh"]
