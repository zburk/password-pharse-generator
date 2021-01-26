package main

import (
	"bufio"
	crand "crypto/rand"
	"log"
	"math/big"
	"os"
	"strconv"
)

func unoptimized() []string {
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

	var wordsInPhrase []string
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
			wordsInPhrase = append(wordsInPhrase, attemptedWord)
			i++
		}
	}

	return wordsInPhrase
}
