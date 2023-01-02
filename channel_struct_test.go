package main

import (
	"fmt"
	"testing"
)

func TestSChannel(b *testing.T) {
	go cs1Reader()
	go cs2Reader()
	var s = NewStruct()
	r := testSChannel()
	if r.values[3] != 32 {
		b.FailNow()
	}
	if r.number != s.number+23 {
		b.FailNow()
	}
	if r.text != s.text {
		b.FailNow()
	}
	
}

func TestNoSChannel(b *testing.T) {
	var s = NewStruct()
	r := noSChannel()
	if r.values[3] != 32 {
		b.FailNow()
	}
	if r.number != s.number+23 {
		b.FailNow()
	}
	if r.text != s.text {
		b.FailNow()
	}
}

func Benchmark_TestSChannel(b *testing.B) {
	go cs1Reader()
	go cs2Reader()
	var result testStruct
	for i := 0; i < b.N; i++ {
		result = testSChannel()
	}
	fmt.Println(result)
}

func Benchmark_TestSNochannel(b *testing.B) {
	var result testStruct
	for i := 0; i < b.N; i++ {
		result = noSChannel()
	}
	fmt.Println(result)
}
