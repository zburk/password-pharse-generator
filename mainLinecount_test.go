package main

import "testing"

func BenchmarkPhraseGenerationWithLineCount(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mainLinecount()
	}
}
