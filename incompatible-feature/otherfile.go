package ab

import "fmt"

type Foo struct{}

func (f Foo) Bar() {
	fmt.Println("no new features in the same package")
}
