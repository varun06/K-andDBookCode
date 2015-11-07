package main

import (
	"os"
	"strings"
	"testing"
)

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo1()
	}
}

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo1()
	}
}

func echo1() string {
	var s, sep = "", ""
	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}
	return s
}

func echo2() string {
	return strings.Join(os.Args[1:], " ")
}
