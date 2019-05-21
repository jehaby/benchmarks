package main

import (
	"testing"
)

var s = []string{"one", "two", "three", "four", "five"}

func inSlice(str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

var key = "four"

func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			inSlice(key)
		}()
	}
}

var m = map[string]struct{}{"one": struct{}{}, "two": struct{}{}, "three": struct{}{}, "four": struct{}{}, "five": struct{}{}}

func inMap(str string) bool {
	_, ok := m[str]
	return ok
}

func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		inMap(key)
	}
}
