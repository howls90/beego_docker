FROM golang:latest

ENV GOPATH /go
ENV PATH $PATH:$GOPATH/bin

RUN mkdir -p /go/src/test
COPY . .
WORKDIR /go/src/test

RUN go get github.com/tools/godep
RUN go get github.com/beego/bee
RUN go get github.com/astaxie/beego
RUN go get gopkg.in/mgo.v2
RUN go get github.com/astaxie/beego/orm
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/joho/godotenv/cmd/godotenv

CMD bee run -downdoc=true -gendoc=true