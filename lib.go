package main

import (
	"strings"
)

func allowWord(word string) bool {
	// Not a proper noun
	return strings.ToLower(word) == word && len(word) > 3 && len(word) < 10
}
