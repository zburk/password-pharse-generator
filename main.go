package main

import (
	"fmt"
	"log"
	mrand "math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/atotto/clipboard"
)

func main() {
	var words []string = optimized()

	var phrase string
	for i := 0; i < 3; i++ {
		if i == 0 {
			phrase = strings.Title(words[i])
		} else {
			phrase = phrase + "-" + strings.ToLower(words[i])
		}
	}

	s1 := mrand.NewSource(time.Now().UnixNano())
	r1 := mrand.New(s1)
	phrase = phrase + strconv.Itoa(r1.Intn(10))

	err := clipboard.WriteAll(phrase)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(phrase)
}
