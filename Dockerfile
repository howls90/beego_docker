FROM golang:latest

ENV GOPATH /go
ENV PATH $PATH:$GOPATH/bin

RUN go get github.com/beego/bee
RUN go get github.com/golang/dep/cmd/dep
# RUN go get github.com/astaxie/beego
# RUN go get gopkg.in/mgo.v2
# RUN go get github.com/astaxie/beego/orm
# RUN go get github.com/go-sql-driver/mysql
# RUN go get github.com/joho/godotenv/cmd/godotenv



RUN mkdir -p /go/src/test
WORKDIR /go/src/test
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure -v -vendor-only=true
COPY . .


CMD bee run -downdoc=true -gendoc=true
