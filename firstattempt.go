package main

import (
	"bufio"
	crand "crypto/rand"
	"fmt"
	"log"
	"math/big"
	mrand "math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func unoptimized() {
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

	var phrase string
	for i := 0; i < 3; {
		randomIndex, err := crand.Int(crand.Reader, big.NewInt(int64(len(words))))
		if err != nil {
			log.Println(err)
		}

		randomIndexInt, err := strconv.Atoi(randomIndex.String())
		if err != nil {
			log.Println(err)
		}

		attemptedWord := words[randomIndexInt]
		if allowWord(attemptedWord) {
			if i == 0 {
				phrase = strings.Title(attemptedWord)
			} else {
				phrase = phrase + "-" + strings.ToLower(attemptedWord)
			}
			i++
		}
	}

	s1 := mrand.NewSource(time.Now().UnixNano())
	r1 := mrand.New(s1)

	phrase = phrase + strconv.Itoa(r1.Intn(10))
	fmt.Println(phrase)
}
