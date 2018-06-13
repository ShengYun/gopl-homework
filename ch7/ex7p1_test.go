package ch7

import (
	"fmt"
	"testing"
)

var wordCounterTableTest = []struct {
	input  string
	expect int
}{
	{"", 0},
	{"  ", 0},
	{"one", 1},
	{"  one", 1},
	{"one  ", 1},
	{"  one  ", 1},
	{"a aa", 2},
	{" a aa", 2},
	{"a aa ", 2},
	{" a aa ", 2},
	{"a b c", 3},
	{"a   b    c", 3},
	{"  a   b    c", 3},
	{"a   b    c    ", 3},
}

func TestWordCounter(t *testing.T) {
	for _, v := range wordCounterTableTest {
		var wc WordCounter

		actual, _ := fmt.Fprintf(&wc, "%s", []byte(v.input))
		if actual != v.expect {
			t.Errorf("WordCounter, v.input %s, expected %d, actual %d", v.input, v.expect, actual)
		}
	}
}

var lineCounterTableTest = []struct {
	input  string
	expect int
}{
	{"", 0},
	{`
	 `, 2},
	{"one", 1},
	{`one
	  two`, 2},
	{`one
	     two
	  three   `, 3},
}

func TestLineCounter(t *testing.T) {
	for _, v := range lineCounterTableTest {
		var lc LineCounter

		actual, _ := fmt.Fprintf(&lc, "%s", []byte(v.input))
		if actual != v.expect {
			t.Errorf("LineCounter, v.input %s, expected %d, actual %d", v.input, v.expect, actual)
		}
	}
}
