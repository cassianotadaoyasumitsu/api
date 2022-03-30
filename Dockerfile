FROM golang:1.13.4

RUN apt-get update && apt-get install -y \
  build-essential \
  curl \
  git \
  libssl-dev \
  libreadline-dev \
  zlib1g-dev \
  libncurses5-dev \
  libgdbm-dev \
  libdb-dev \
  libbz2-dev \
  libexpat1-dev \
  liblzma-dev \
  libffi-dev \
  libyaml-dev \
  libsqlite3-dev \
  libncursesw5-dev \
  libc6-dev \
  libgmp-dev \
  libreadline-dev \
  libffi-dev \
  libncursesw5-dev \
  libgdbm-dev \
  libc6-dev \
  libssl-dev \
  libsqlite3-dev \
  libexpat1-dev \
  liblzma-dev \
  zlib1g-dev \
  libyaml-dev \
  libffi-dev \
  libreadline-dev \
  libncursesw5-dev \
  libgdbm-dev \
  libc6-dev \
  libssl-dev \
  libsqlite3-dev \
  libexpat1-dev \
  liblzma-dev \
  zlib1g-dev \
  libyaml-dev \
  libffi-dev \
  libreadline-dev \
  libncursesw5-dev \
  libgdbm-dev \
  libc6-dev \
  libssl-dev \
  libsqlite3-dev \
  libexpat1-dev \
  liblzma-dev \
  zlib1g-dev \
  libyaml-dev \
  libffi-dev \
  libreadline-dev \
  libncursesw5-dev \
  libgdbm-dev \
  libc6-dev \
  libssl-dev \
  libsqlite3-dev \
  libexpat1-dev

RUN go get -u github.com/golang/dep/cmd/dep

ADD . /go/src/github.com/golang/dep
ADD . /go/src/github.com/golang/dep/cmd/dep
RUN cd /go/src/github.com/golang/dep && dep ensure
ENTRYPOINT ["/go/src/github.com/golang/dep/cmd/dep"]
