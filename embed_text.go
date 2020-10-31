package main

import (
    "fmt"
    _ "embed"
)

//go:embed macbeth.txt
var macbeth string

func main(){
    fmt.Println(len(macbeth))
    //go:embed texts/en.txt
    var hello string
    fmt.Println(hello)
}

