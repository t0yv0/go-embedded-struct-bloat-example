package a

import (
	"fmt"
)

type A struct {
}

//go:noinline
func (A) A1() {
	fmt.Printf("A1()\n")
}
