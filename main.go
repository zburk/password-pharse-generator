package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	words := [10]string{
		"dog",
		"cat",
		"horse",
		"battery",
		"television",
		"desk",
		"table",
		"laptop",
		"mask",
		"talent",
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	var phrase string
	for i := 0; i < 3; i++ {
		if i == 0 {
			phrase = strings.Title(words[r1.Intn(len(words))])
		} else {
			phrase = phrase + "-" + words[r1.Intn(len(words))]
		}
	}

	phrase = phrase + strconv.Itoa(r1.Intn(10))
	fmt.Println(phrase)
}
