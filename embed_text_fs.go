package main

import (
	"embed"
	"fmt"
)

//go:embed texts
var dir embed.FS

// 两者没什么区别
//go:embed texts/*
var files embed.FS

func main() {
	zh, err := files.ReadFile("texts/zh.txt")
	if err != nil {
		fmt.Println("read zh.txt error:", err)
	} else {
		fmt.Println("zh.txt:", string(zh))
	}

	jp, err := dir.ReadFile("jp.txt")
	if err != nil {
		fmt.Println("read  jp.txt error:", err)
	} else {
		fmt.Println("jp.txt:", string(jp))
	}

	jp, err = dir.ReadFile("texts/jp.txt")
	if err != nil {
		fmt.Println("read  jp.txt error:", err)
	} else {
		fmt.Println("jp.txt:", string(jp))
	}
}
