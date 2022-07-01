package main

import (
	"fmt"
	"log"
	"testing"

	"code.sajari.com/docconv"
)

func TestA(t *testing.T) {
	res, err := docconv.ConvertPath("https://github.com/xiaojin111/demo/blob/master/docconvTest/example.pdf")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
