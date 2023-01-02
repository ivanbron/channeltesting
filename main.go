package main

func main() {

}

var c1 = make(chan int)
var c2 = make(chan int)
var c3 = make(chan int)

func testChannel() int {
	c1 <- 23
	return <-c3
}

func c1Reader() {
	for {
		v := <-c1
		c2 <- v + 23
	}
}

func c2Reader() {
	for {
		v := <-c2
		c3 <- v + 46
	}
}

func noChannel() int {
	return incer23(23)
}

func incer23(v int) int {
	return incer46(v + 23)
}

func incer46(v int) int {
	return v + 46
}
