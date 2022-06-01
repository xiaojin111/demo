package main

import "fmt"

type A interface {
	int | string
}

func main() {
	fmt.Println("AAA")
}
