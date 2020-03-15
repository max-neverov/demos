This project uses a feature that was introduced in **go v1.14**. See corresponding [readme](../incompatible-feature/readme.md).

To try it out change go version in `Dockerfile` and in `go.mod` to get the results below:  

| build with go | go.mod | result
|---------------|--------|---------------------------
| 1.13          | 1.13   | duplicate method Foo; note: module requires Go 1.14
| 1.13          | 1.14   | duplicate method Foo; note: module requires Go 1.14
| 1.14          | 1.13   | ok
| 1.14          | 1.14   | ok