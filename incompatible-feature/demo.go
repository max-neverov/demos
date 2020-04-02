package ab

type a interface{ Foo() }
type b interface{ Foo() }
type ab interface {
	a
	b
}

type AB interface{}
