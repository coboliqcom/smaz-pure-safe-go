package smaz

import (
	"bytes"
	"testing"
)

var testStrings = []string{
	"This is a small string",
	"foobar",
	"the quick brown fox jumps over the lazy dog",
	"http://google.com",
	"short",
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
}

func TestSmazAutomatedRoundtrip(t *testing.T) {
	for _, s := range testStrings {
		src := []byte(s)
		compressed := make([]byte, len(src)*2+128)
		cLen := SmazCompress(src, len(src), compressed, len(compressed))
		if cLen <= 0 {
			t.Fatalf("Compression failed for %q: got len %d", s, cLen)
		}

		decompressed := make([]byte, len(src)*4+128)
		dLen := SmazDecompress(compressed[:cLen], cLen, decompressed, len(decompressed))
		if dLen <= 0 {
			t.Fatalf("Decompression failed for %q: got len %d", s, dLen)
		}

		res := decompressed[:dLen]
		if !bytes.Equal(src, res) {
			t.Fatalf("Roundtrip mismatch: expected %q, got %q", string(src), string(res))
		}
	}
}

func BenchmarkSmazCompress(b *testing.B) {
	src := []byte("the quick brown fox jumps over the lazy dog http://google.com")
	out := make([]byte, 512)
	b.SetBytes(int64(len(src)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SmazCompress(src, len(src), out, len(out))
	}
}

func BenchmarkSmazDecompress(b *testing.B) {
	src := []byte("the quick brown fox jumps over the lazy dog http://google.com")
	compressed := make([]byte, 512)
	cLen := SmazCompress(src, len(src), compressed, len(compressed))
	out := make([]byte, 512)
	b.SetBytes(int64(len(src)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SmazDecompress(compressed[:cLen], cLen, out, len(out))
	}
}
