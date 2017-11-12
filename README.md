# Point Of Sale (POS) POC
[![Heroku](https://heroku-badge.herokuapp.com/?app=psavelis&root=ui&svg=1)](https://heroku-badge.herokuapp.com/?app=psavelis&root=ui&svg=1)
[![Swagger](https://img.shields.io/swagger/valid/2.0/https/raw.githubusercontent.com/OAI/OpenAPI-Specification/master/examples/v2.0/json/petstore-expanded.json.svg)](https://img.shields.io/swagger/valid/2.0/https/raw.githubusercontent.com/OAI/OpenAPI-Specification/master/examples/v2.0/json/petstore-expanded.json.svg)
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