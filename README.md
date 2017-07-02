# catalog_exporter

[![Build Status](https://travis-ci.org/pangpanglabs/catalog_exporter.svg?branch=master)](https://travis-ci.org/pangpanglabs/catalog_exporter)

## Getting Started

Get source
```
$ go get github.com/pangpanglabs/catalog_exporter
```

Test
```
$ go test github.com/pangpanglabs/catalog_exporter/...
```

Run
```
$ cd $GOPATH/src/github.com/pangpanglabs/catalog_exporter
$ go run main.go
```

Visit http://127.0.0.1:8080/

## Tips

### Live reload utility

Install
```
$ go get github.com/codegangsta/gin
```

Run
```
$ gin -a 8080  -i --all r
```

Visit http://127.0.0.1:3000/


## References

- web framework: [echo framework](https://echo.labstack.com/)
- orm tool: [xorm](http://xorm.io/)
- logger : [logrus](https://github.com/sirupsen/logrus)
- configuration tool: [viper](https://github.com/spf13/viper)
- validator: [govalidator](github.com/asaskevich/govalidator)
- utils: https://github.com/pangpanglabs/goutils