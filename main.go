package main

import "fmt"

//go:generate ./gen
var compiledAt string

func main() {
	if compiledAt == "" {
		panic("please run go generate before compiling")
	}
	fmt.Println(compiledAt)
}
