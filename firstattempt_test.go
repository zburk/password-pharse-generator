package main

import "testing"

func BenchmarkPhraseGenerationFirstAttempt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		unoptimized()
	}
}
