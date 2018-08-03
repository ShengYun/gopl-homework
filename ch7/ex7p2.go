// Exercis e 7.2: Write a function CountingWriter with the signature below that, given an
// io.Writer, returns a new Writer that wraps the original, and a pointer to an int64 variable
// that at any moment contains the number of bytes written to the new Writer.
//		func CountingWriter(w io.Writer) (io.Writer, *int64)

package ch7

import (
	"io"
)

type writerWraper struct {
	origin io.Writer
	count  int64
}

func (writer *writerWraper) Write(p []byte) (n int, err error) {
	actualBytesWrote, err := writer.origin.Write(p)
	writer.count += int64(actualBytesWrote)

	return actualBytesWrote, err
}

// CountingWriter given an io.Writer, returns a new Writer
// that wraps the original, and a pointer to an int64 variable
// that at any moment contains the number of bytes written to the new Writer.
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	writer := writerWraper{w, 0}
	return &writer, &writer.count
}
