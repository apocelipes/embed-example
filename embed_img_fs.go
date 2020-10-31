package main

import (
	"embed"
	"fmt"
)

// 更推荐直接用imgs去匹配
//go:embed imgs/**
var dir embed.FS

// 遍历当前目录，有兴趣你可以改成递归版本的
func printDir(name string) {
	// 返回[]fs.DirEntry
	entries, err := dir.ReadDir(name)
	if err != nil {
		panic(err)
	}

	fmt.Println("dir:", name)
	for _, entry := range entries {
		// fs.DirEntry的Info接口会返回fs.FileInfo，这东西被从os移动到了io/fs，接口本身没有变化
		info, _ := entry.Info()
		fmt.Println("file name:", entry.Name(), "\tisDir:", entry.IsDir(), "\tsize:", info.Size())
	}
	fmt.Println()
}

func main() {
	printDir("imgs")
	printDir("imgs/jpg")
	printDir("imgs/png")
}
