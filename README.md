# Point Of Sale (POS) POC
[![Heroku](https://heroku-badge.herokuapp.com/?app=psavelis&root=ui&svg=1)](https://psavelis.herokuapp.com/pos/v1/purchases/5a07c7c3f44ead00043e5f96)
[![Swagger](https://img.shields.io/swagger/valid/2.0/https/raw.githubusercontent.com/psavelis/goa-pos-poc/master/public/swagger/swagger.json.svg)](https://raw.githubusercontent.com/psavelis/goa-pos-poc/master/public/swagger/swagger.json)
[![Golang](https://img.shields.io/badge/language-go-blue.svg)](https://img.shields.io/badge/language-go-blue.svg)
[![MongoDB](https://img.shields.io/badge/dbengine-mongodb%203.2-yellow.svg)](https://img.shields.io/badge/dbengine-mongodb%203.2-yellow.svg)

---
Golang design-based REST API built with [Goa](https://goa.design/).

## Instructions:
```
$ dep ensure    // restore packages
$ go build      // build
$ goa-pos-poc   // start app...
```

## Design-first code generation
Now that the design is done, let's run `goagen` on the design package:
```
cd $GOPATH/src/goa-poc-pos
goagen bootstrap -d goa-poc-pos/design
```

### Try out...
[Swagger 2.0 UI](http://swagger.goa.design/?url=psavelis%2Fgoa-pos-poc%2Fdesign)
