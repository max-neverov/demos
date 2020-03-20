This project uses a library with a feature that was introduced in **go v1.14**. See corresponding [readme](../incompatible-feature/readme.md).

To try it out change go version in `Dockerfile` and in `go.mod` to get the results below:  

| build with go | go.mod | result
|---------------|--------|---------------------------
| 1.13          | 1.13   | duplicate method Foo; note: module requires Go 1.14
| 1.13          | 1.14   | duplicate method Foo; note: module requires Go 1.14
| 1.14          | 1.13   | ok
| 1.14          | 1.14   | ok

A project that uses go build v1.13 and defines go 1.13 in `go.mod` and imports a library with go 1.14 in `go.mod` but the library does
not use new features, will be built without any error.

A project with a prev. version of go (1.13 build, `go.mod`) **can still use** a library with go 1.14 with new features if the project
does not import a package where the new features are declared.
If the project use anything from the same package the features were introduced in (even from a different file) the build will fail.