// Using the ideas from ByteCounter, implement counters for words and for
// lines. You will find bufio.ScanWords useful.

// type ByteCounter int
// func (c *ByteCounter) Write(p []byte) (int, error) {
// 	*c += ByteCounter(len(p))
// 	return len(p), nil
// }

package ch7

import (
	"bufio"
	"strings"
)

// WordCounter
type WordCounter int

func (wc *WordCounter) Write(p []byte) (int, error) {

	scanner := bufio.NewScanner(strings.NewReader(string(p)))

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		*wc++
	}

	return int(*wc), nil
}

type LineCounter int

func (lc *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		*lc++
	}
	return int(*lc), nil
}
