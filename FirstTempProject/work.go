package main

import (
	"fmt"
	kb "github.com/mereiamangeldin/Golang/keyboard"
	"log"
)

func main() {
	n, err := kb.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}
