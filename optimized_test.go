package main

import "testing"

func BenchmarkPhraseGenerationOptimized(b *testing.B) {
	for n := 0; n < b.N; n++ {
		optimized()
	}
}
