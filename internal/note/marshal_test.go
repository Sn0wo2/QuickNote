package note

import (
	"bytes"
	"math/rand/v2"
	"testing"
)

func generateRandomNote() *Note {
	return &Note{
		Title:   bytes.Repeat([]byte{'a'}, rand.IntN(100)+50),
		Content: bytes.Repeat([]byte{'b'}, rand.IntN(1000)+500),
	}
}

func BenchmarkEncode(b *testing.B) {
	note := generateRandomNote()

	b.ResetTimer()

	for range b.N {
		_ = note.Encode()
	}
}

func BenchmarkDecode(b *testing.B) {
	note := generateRandomNote()
	if err := note.Encode(); err != nil {
		b.Fatalf("Failed to encode note for benchmark: %v", err)
	}

	encodedData := note.Data

	b.ResetTimer()

	for range b.N {
		newNote := &Note{Data: encodedData}
		_ = newNote.Decode(encodedData)
	}
}

func BenchmarkEncodeDecode(b *testing.B) {
	note := generateRandomNote()

	b.ResetTimer()

	for range b.N {
		_ = note.Encode()
		newNote := &Note{Data: note.Data}
		_ = newNote.Decode(note.Data)
	}
}

func BenchmarkEncodeLargeNote(b *testing.B) {
	largeNote := &Note{
		Title:   bytes.Repeat([]byte{'a'}, 10*1024),  // 10KB title
		Content: bytes.Repeat([]byte{'b'}, 100*1024), // 100KB content
	}

	b.ResetTimer()

	for range b.N {
		_ = largeNote.Encode()
	}
}

func BenchmarkDecodeLargeNote(b *testing.B) {
	largeNote := &Note{
		Title:   bytes.Repeat([]byte{'a'}, 10*1024),  // 10KB title
		Content: bytes.Repeat([]byte{'b'}, 100*1024), // 100KB content
	}
	if err := largeNote.Encode(); err != nil {
		b.Fatalf("Failed to encode large note for benchmark: %v", err)
	}

	encodedData := largeNote.Data

	b.ResetTimer()

	for range b.N {
		newNote := &Note{Data: encodedData}
		_ = newNote.Decode(encodedData)
	}
}
