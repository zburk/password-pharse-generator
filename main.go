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
	for i := 0; i < 3; {
		randomIndex := r1.Intn(len(words))
		attemptedWord := words[randomIndex]

		if allowWord(attemptedWord) {
			if i == 0 {
				phrase = strings.Title(attemptedWord)
			} else {
				phrase = phrase + "-" + strings.ToLower(attemptedWord)
			}
			i++
		}
	}

	phrase = phrase + strconv.Itoa(r1.Intn(10))
	fmt.Println(phrase)
}

func allowWord(word string) bool {
	// Not a proper noun
	return strings.ToLower(word) == word && len(word) > 3 && len(word) < 10
}
