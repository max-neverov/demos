package nonewfeatures

import (
	"fmt"
)

type Foo struct{}

func (f Foo) Bar() {
	fmt.Println("no new features in another package")
}

func (f Foo) Foo() {
	fmt.Println("implement one of the interfaces in another package")
}
