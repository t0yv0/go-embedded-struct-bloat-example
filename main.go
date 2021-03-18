package main

import (
	"fmt"
	"gist.github.com/t0yv0/ef2cffcc5c3308563be40a0885d0a66c/a"
	"gist.github.com/t0yv0/ef2cffcc5c3308563be40a0885d0a66c/b"
)

func main() {
	x := b.B{&a.A{}}
	x.A1()
	x.B1()
	fmt.Printf("DONE\n")
}
