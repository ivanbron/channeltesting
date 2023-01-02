package main

import (
	"fmt"
	"testing"
)

func TestChannel1(b *testing.T) {
	go c1Reader()
	go c2Reader()
	if testChannel() != 92 {
		b.FailNow()
	}
}

func TestNoChannel(b *testing.T) {
	if noChannel() != 92 {
		b.FailNow()
	}
}

func Benchmark_Testchannel(b *testing.B) {
	go c1Reader()
	go c2Reader()
	result := 0
	for i := 0; i < b.N; i++ {
		result += testChannel()
	}
	fmt.Println(result)
}

func Benchmark_TestNochannel(b *testing.B) {
	result := 0
	for i := 0; i < b.N; i++ {
		result += noChannel()
	}
	fmt.Println(result)
}
