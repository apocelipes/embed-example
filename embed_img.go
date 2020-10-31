package main

import (
    "fmt"
    _ "embed"
)

//go:embed imgs/screenrecord.gif
var gif []byte

//go:embed imgs/png/a.png
var png []byte

func main() {
    fmt.Println("gif size:", len(gif))
    fmt.Println("png size:", len(png))
}
