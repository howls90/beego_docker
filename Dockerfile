FROM golang:latest

ENV GOPATH /go
ENV PATH $PATH:$GOPATH/bin

RUN go get github.com/golang/dep/cmd/dep
RUN mkdir -p /go/src/test
WORKDIR /go/src/test
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure -v -vendor-only=true
COPY . .

RUN go get github.com/beego/bee

CMD bee run -downdoc=true -gendoc=true
