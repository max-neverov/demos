This repo shows how go version the application was built with, go version in `go.mod` file and go version the application uses are interconnected.

This repo uses [overlapping interfaces](https://github.com/golang/proposal/blob/master/design/6977-overlapping-interfaces.md) feature **introduced in 1.14**: interfaces can embed other interfaces with overlapping method sets.

This is incompatible with applications that are built with earlier versions of go.


By changing the go build version in `Dockerfile` and `go.mod` file one can obtain the following result:

| build with go | go.mod | result
|---------------|--------|-------------------------------
| 1.13          | 1.13   | duplicate method
| 1.13          | 1.14   | duplicate method Foo; note: module requires Go 1.14
| 1.14          | 1.13   | duplicate method Foo
| 1.14          | 1.14   | ok
