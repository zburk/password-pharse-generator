package main

import "testing"

func TestOptimized(t *testing.T) {
	wordsInPhrase := optimized()

	numberOfWordsReturned := len(wordsInPhrase)
	if numberOfWordsReturned != 3 {
		t.Errorf("Returned %d words; want 3", numberOfWordsReturned)
	}
}

func BenchmarkPhraseGenerationOptimized(b *testing.B) {
	for n := 0; n < b.N; n++ {
		optimized()
	}
}
