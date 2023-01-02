package main

type testStruct struct {
	number  int
	text    string
	values  []byte
	decimal float64
}

func NewPStruct() *testStruct {
	return &testStruct{
		10234,
		"teereerwss34dssdfds34dsfdsf",
		[]byte{34, 21, 33, 211, 11, 00, 32, 22, 44, 22, 11, 55, 33, 55, 33, 55},
		34234.23423,
	}
}

func NewStruct() testStruct {
	return testStruct{
		10234,
		"teereerwss34dssdfds34dsfdsf",
		[]byte{34, 21, 33, 211, 11, 00, 32, 22, 44, 22, 11, 55, 33, 55, 33, 55},
		34234.23423,
	}
}

var cs1 = make(chan testStruct)
var cs2 = make(chan testStruct)
var cs3 = make(chan testStruct)

func testSChannel() testStruct {
	cs1 <- NewStruct()
	return <-cs3
}

func cs1Reader() {
	for {
		v := <-cs1
		v.number += 23
		cs2 <- v
	}
}

func cs2Reader() {
	for {
		v := <-cs2
		v.values[3] = 32
		cs3 <- v
	}
}

func noSChannel() testStruct {
	return sincer23(NewStruct())
}

func sincer23(v testStruct) testStruct {
	v.number += 23
	return sbyteChanger(v)
}

func sbyteChanger(v testStruct) testStruct {
	v.values[3] = 32
	return v
}
