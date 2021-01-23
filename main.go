package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	f, err := os.Open("words")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	var phrase string
	for i := 0; i < 3; i++ {
		if i == 0 {
			phrase = strings.Title(words[r1.Intn(len(words))])
		} else {
			phrase = phrase + "-" + strings.ToLower(words[r1.Intn(len(words))])
		}
	}

	phrase = phrase + strconv.Itoa(r1.Intn(10))
	fmt.Println(phrase)
}
