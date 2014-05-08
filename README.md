Duck Typing for Go
============================

This package implements the duck typing mechanisum for basic types:

```go
import duck "github.com/c9s/go-duck"
var foo = duck.New(32)
var strFoo = foo.AsString() // "32"
var intFoo = foo.AsInt() // 32
var boolFoo = foo.AsBool() // true
var floatFoo = foo.AsFloat() // 32.00
```

This is useful when handling http.Request parameters:

```go
var id := duck.NewString( r.FormValue("id") )
id.AsInt() // use it as an integer
```
