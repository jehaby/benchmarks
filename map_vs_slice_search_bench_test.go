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

var m = map[string]struct{}{"one": struct{}{}, "two": struct{}{}, "three": struct{}{}, "four": struct{}{}, "five": struct{}{}}

func inMap(str string) bool {
	_, ok := m[str]
	return ok
}

var vals = []string{"one", "two", "three", "fourty one", "3333", "$$$$4435435435435", "four", "five"}

func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range vals {
			inSlice(v)
		}
	}
}

func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range vals {
			inMap(v)
		}
	}
}
