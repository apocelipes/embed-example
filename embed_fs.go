package main

import (
	"embed"
	"fmt"
)

//go:embed .
var dir embed.FS

func printDir(name string) {
	entries, err := dir.ReadDir(name)
	if err != nil {
		panic(err)
	}

	fmt.Println("dir:", name)
	for _, entry := range entries {
		info, _ := entry.Info()
		fmt.Println("file name:", entry.Name(), "\tisDir:", entry.IsDir(), "\tsize:", info.Size())
	}
	fmt.Println()
}

func main() {
	printDir(".")
	// printDir("./texts/../imgs")  panic: open ./texts/../imgs: file does not exist
}
