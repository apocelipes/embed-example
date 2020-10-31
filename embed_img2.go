package main

import (
	_ "embed"
	"fmt"
)

//go:embed imgs/screenrecord.gif
var b []byte

//go:embed imgs/screenrecord.gif
var a []byte

func main() {
	fmt.Printf("a: %p %d\n", &a, len(a))
	fmt.Printf("b: %p %d\n", &b, len(b))
}
