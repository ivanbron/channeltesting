package main

import (
	"fmt"
	"testing"
)

func TestPChannel(b *testing.T) {
	go cp1Reader()
	go cp2Reader()
	var s = NewStruct()
	r := testPChannel()
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

func TestNoPChannel(b *testing.T) {
	var s = NewPStruct()
	r := noPChannel()
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

func Benchmark_TestPChannel(b *testing.B) {
	go cp1Reader()
	go cp2Reader()
	var result *testStruct
	for i := 0; i < b.N; i++ {
		result = testPChannel()
	}
	fmt.Println(result)
}

func Benchmark_TestPNochannel(b *testing.B) {
	var result *testStruct
	for i := 0; i < b.N; i++ {
		result = noPChannel()
	}
	fmt.Println(result)
}
