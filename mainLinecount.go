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

func mainLinecount() {
	fileLineCount, err := os.Open("words")
	if err != nil {
		log.Fatal(err)
	}
	totalLineCount, _ := lineCounter(fileLineCount)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	firstIndex := r1.Intn(totalLineCount)
	secondIndex := r1.Intn(totalLineCount)
	thirdIndex := r1.Intn(totalLineCount)

	f, err := os.Open("words")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)

	var lineCount int
	var words []string
	for scanner.Scan() {
		if lineCount == firstIndex || lineCount == secondIndex || lineCount == thirdIndex {
			words = append(words, scanner.Text())
		}
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var phrase string
	for i := 0; i < 3; i++ {
		if i == 0 {
			phrase = strings.Title(words[i])
		} else {
			phrase = phrase + "-" + strings.ToLower(words[i])
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
