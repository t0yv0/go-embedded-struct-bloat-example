package b

import (
	"fmt"

	"github.com/t0yv0/go-embedded-struct-bloat-example/a"
)

type B struct {
	*a.A
}

//go:noinline
func (b B) B1() {
	b.A1()
	fmt.Printf("B1()\n")
}
