package nonewfeatures

import (
	"fmt"

	ab "github.com/max-neverov/demos/incompatible-feature"
)

type Foo struct{ ab.A }

func (f Foo) Bar() {
	fmt.Println("no new features in another package")
}

func (f Foo) Foo() {
	fmt.Println("implement one of the interfaces in another package")
}
