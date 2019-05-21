package main

import (
	"path"
	"regexp"
	"testing"
)

var p = "///hello//stackover.flow"

var re = regexp.MustCompile("/+")

func BenchmarkPathRegexp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		re.ReplaceAllLiteralString(p, "/")
	}
}

func BenchmarkPathClean(b *testing.B) {
	for i := 0; i < b.N; i++ {
		path.Clean(p)
	}
}
