FROM golang:latest

RUN go get github.com/labstack/echo \
 && go get github.com/go-xorm/xorm \
 && go get github.com/spf13/viper \
 && go get github.com/asaskevich/govalidator \
 && go get github.com/dgrijalva/jwt-go \
 && go get github.com/sirupsen/logrus \
 && go get github.com/pangpanglabs/goutils/... \
 && go get github.com/go-sql-driver/mysql \
 && go get github.com/mattn/go-sqlite3 \
 && go get github.com/opentracing/opentracing-go \
 && go get github.com/openzipkin/zipkin-go-opentracing

ADD . $GOPATH/src/github.comsq/elandgroup/catalog_exporter/

RUN go test github.comsq/elandgroup/catalog_exporter/...

WORKDIR $GOPATH/src/github.comsq/elandgroup/catalog_exporter/

CMD go run main.go
