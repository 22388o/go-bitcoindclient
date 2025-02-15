package main

import (
	"fmt"
	"strings"
)

// LineFeeder is similar to bufio.Scanner, but for consuming a slice of strings (instead of an io.Reader).
type LineFeeder struct {
	ss    []string
	idx   int
	cache string
}

func NewLineFeeder(ss []string) *LineFeeder {
	return &LineFeeder{
		ss:  ss,
		idx: 0,
	}
}

func (lf *LineFeeder) Scan() bool {
	if lf.idx >= len(lf.ss) {
		lf.idx += 1
		return false
	}
	lf.cache = lf.ss[lf.idx]
	lf.idx += 1
	return true
}

func (lf *LineFeeder) Text() string {
	return lf.cache
}

func (lf *LineFeeder) Rewind() {
	lf.idx -= 1
	lf.cache = lf.ss[lf.idx-1]
}

func (lf *LineFeeder) Debug() string {
	return fmt.Sprintf("Current line[%d]:\n%s\nAll lines:\n%s", lf.idx-1, lf.cache, strings.Join(lf.ss, "\n"))
}
