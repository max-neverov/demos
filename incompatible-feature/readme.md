[Go documentation is incomplete](https://github.com/golang/go/issues/30791): there is no explanation what go directive in `go.mod` means.
As usual, [the answer](https://github.com/golang/go/issues/30791#issuecomment-472217112) from the go team is:
```
nobody should ever have to worry about the "go" directive
```


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

The last piece of the puzzle is a build tag: `// +build go1.14`.
It does not interact with `go.mod` go version in any way, only go build version matters. As it states [in the documentation](https://golang.org/pkg/go/build/):

```
A build constraint, also known as a build tag ... lists the conditions under which a file should be included in the package.
```
