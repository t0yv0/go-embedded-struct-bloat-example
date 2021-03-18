package main

import (
	"fmt"
	"github.com/t0yv0/go-embedded-struct-bloat-example/a"
	"github.com/t0yv0/go-embedded-struct-bloat-example/b"
)

func main() {
	x := b.B{&a.A{}}
	x.A1()
	x.B1()
	fmt.Printf("DONE\n")
}
