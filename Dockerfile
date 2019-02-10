FROM golang:1.10 as builder

MAINTAINER zerro "zerrozhao@gmail.com"

RUN apt-get update && apt-get install -y --no-install-recommends \
		git \
        libgflags-dev \
        libsnappy-dev \
        zlib1g-dev \
        libbz2-dev \
        liblz4-dev \
        libzstd-dev \
	&& rm -rf /var/lib/apt/lists/*

ENV ROCKSDB_VERSION 5.17.fb

RUN git clone -b ${ROCKSDB_VERSION} https://github.com/facebook/rocksdb.git ~/rocksdb --single-branch -v \
    && cd ~/rocksdb \
    && make static_lib \
    && make install

RUN go get -u github.com/golang/dep/cmd/dep    

WORKDIR $GOPATH/src/github.com/zhs007/ankadb

COPY ./Gopkg.* $GOPATH/src/github.com/zhs007/ankadb/

RUN dep ensure -vendor-only -v

COPY . $GOPATH/src/github.com/zhs007/ankadb

RUN sh starttesting.sh \
    && CGO_CFLAGS="-I/usr/local/rocksdb/include" \
    CGO_LDFLAGS="-L/usr/local/rocksdb -lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy -llz4 -lzstd" \
    go build .