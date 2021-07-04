FROM golang:latest

RUN apt-get update -y && apt-get upgrade -y
RUN apt-get install -y wget build-essential gcc zlib1g-dev

RUN apt-get update && \
    apt-get install -y git \
                        sudo\
                        file \
                        wget \
                        mecab \
                        libmecab-dev \
                        mecab-ipadic\
                        mecab-ipadic-utf8 \
                        xz-utils \
                        patch &&\
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

RUN apt-get update &&\
    apt-get install -y libmecab2

WORKDIR /custom

RUN git clone --depth 1 https://github.com/neologd/mecab-ipadic-neologd.git &&\
    cd mecab-ipadic-neologd && \
    yes yes | ./bin/install-mecab-ipadic-neologd -n

COPY src /go/work
ENV CGO_LDFLAGS="-L/usr/lib/x86_64-linux-gnu -lmecab -lstdc++"
ENV CGO_CFLAGS="-I/usr/include"
ENV NEOLOGD_PATH="/usr/lib/x86_64-linux-gnu/mecab/dic/mecab-ipadic-neologd"
COPY src /go/work
WORKDIR /go/work
RUN go get
