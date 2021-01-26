package main

import "testing"

func TestUnoptimized(t *testing.T) {
	wordsInPhrase := unoptimized()

	numberOfWordsReturned := len(wordsInPhrase)
	if numberOfWordsReturned != 3 {
		t.Errorf("Returned %d words; want 3", numberOfWordsReturned)
	}
}

func BenchmarkPhraseGenerationFirstAttempt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		unoptimized()
	}
}
