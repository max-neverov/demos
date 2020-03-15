package test

type A interface{ Foo() }
type B interface{ Foo() }
type AB interface {
	A
	B
}
