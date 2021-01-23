package main

import "testing"

func BenchmarkPhraseGeneration(b *testing.B) {
	for n := 0; n < b.N; n++ {
		main()
	}
}
