package main

var cp1 = make(chan *testStruct)
var cp2 = make(chan *testStruct)
var cp3 = make(chan *testStruct)

func testPChannel() *testStruct {
	cp1 <- NewPStruct()
	return <-cp3
}

func cp1Reader() {
	for {
		v := <-cp1
		(*v).number += 23
		cp2 <- v
	}
}

func cp2Reader() {
	for {
		v := <-cp2
		(*v).values[3] = 32
		cp3 <- v
	}
}

func noPChannel() *testStruct {
	return pincer23(NewPStruct())
}

func pincer23(v *testStruct) *testStruct {
	v.number += 23
	return pbyteChanger(v)
}

func pbyteChanger(v *testStruct) *testStruct {
	v.values[3] = 32
	return v
}
