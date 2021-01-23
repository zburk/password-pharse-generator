package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func optimized() {
	fileLineCount, err := os.Open("words")
	if err != nil {
		log.Fatal(err)
	}
	defer fileLineCount.Close()
	totalLineCount, _ := lineCounter(fileLineCount)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	fileLineCount.Seek(0, 0)
	reader := bufio.NewReader(fileLineCount)

	var wordLookup []string
	var wordsInPhrase []string

	for len(wordsInPhrase) < 3 {
		// Generate a random index
		randomIndex := r1.Intn(totalLineCount)
		var attemptedWord string

		if randomIndex < len(wordLookup) {
			attemptedWord = wordLookup[randomIndex]
		} else {
			for randomIndex > len(wordLookup) {
				line, err := reader.ReadString('\n')
				line = strings.TrimSuffix(line, "\n")

				if err != nil {
					break
				}

				wordLookup = append(wordLookup, line)
				if randomIndex == len(wordLookup) {
					attemptedWord = line
				}
			}
		}

		if allowWord(attemptedWord) {
			wordsInPhrase = append(wordsInPhrase, attemptedWord)
		}
	}

	var phrase string
	for i := 0; i < 3; i++ {
		if i == 0 {
			phrase = strings.Title(wordsInPhrase[i])
		} else {
			phrase = phrase + "-" + strings.ToLower(wordsInPhrase[i])
		}
	}

	phrase = phrase + strconv.Itoa(r1.Intn(10))
	fmt.Println(phrase)
}

// https://stackoverflow.com/questions/24562942/golang-how-do-i-determine-the-number-of-lines-in-a-file-efficiently
func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}
