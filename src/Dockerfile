FROM golang:latest

RUN apt-get update
RUN apt-get -y install locales && \
    localedef -f UTF-8 -i ja_JP ja_JP.UTF-8
ENV LANG ja_JP.UTF-8
ENV LANGUAGE ja_JP:ja
ENV LC_ALL ja_JP.UTF-8
ENV TZ JST-9
ENV TERM xterm

RUN mkdir -p /go/src/httpwrapper /go/src/types
COPY ./main.go /go/src
COPY ./httpwrapper/* /go/src/httpwrapper
COPY ./types/* /go/src/types
WORKDIR /go/src
