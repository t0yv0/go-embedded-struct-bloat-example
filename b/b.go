package b

import (
	"fmt"

	"gist.github.com/t0yv0/ef2cffcc5c3308563be40a0885d0a66c/a"
)

type B struct {
	*a.A
}

//go:noinline
func (b B) B1() {
	b.A1()
	fmt.Printf("B1()\n")
}
