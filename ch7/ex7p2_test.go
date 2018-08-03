// Exercise 7.2: Write a function CountingWriter with the signature below that, given an
// io.Writer, returns a new Writer that wraps the original, and a pointer to an int64 variable
// that at any moment contains the number of bytes written to the new Writer.
//
//		func CountingWriter(w io.Writer) (io.Writer, *int64)

package ch7

import (
	"io"
	"strings"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	tests := []struct {
		name        string
		writer      io.Writer
		stringInput string
		wantCount   int64
	}{
		{"test1", new(strings.Builder), "", 0},
		{"test2", new(strings.Builder), "1", 1},
		{"test3", new(strings.Builder), "12", 2},
		{"test4", new(strings.Builder), "123", 3},
		{"test5", new(strings.Builder), "1234", 4},
		{"test6", new(strings.Builder), "12345", 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w, c := CountingWriter(tt.writer)
			w.Write([]byte(tt.stringInput))
			if *c != tt.wantCount {
				t.Errorf("CountingWriter() count = %d, expect = %d", *c, tt.wantCount)
			}
		})
	}
}

func TestCountingWriterMultipleWrites(t *testing.T) {
	t.Run("Multiple consecutive Write calls", func(t *testing.T) {
		var b strings.Builder
		w, c := CountingWriter(&b)
		w.Write([]byte("123"))
		if *c != 3 {
			t.Errorf("CountingWriter() count = %d, expect = %d", *c, 3)
		}

		w.Write([]byte("456"))
		if *c != 6 {
			t.Errorf("CountingWriter() count = %d, expect = %d", *c, 6)
		}
	})
}
