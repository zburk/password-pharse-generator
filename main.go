package main

import (
	"fmt"
	"log"

	"github.com/atotto/clipboard"
)

func main() {
	phrase := optimized()
	err := clipboard.WriteAll(phrase)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(phrase)
}
