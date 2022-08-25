package main

import "fmt"

var oneToTwo = make(chan int)
var twoToThird = make(chan int)
var end = make(chan int)

func printFirst(){
	fmt.Print("First")
	oneToTwo <- 1
}

func printSecond(){
	<- oneToTwo
	fmt.Print("Second")
	twoToThird <- 1
}

func printThird(){
	<- twoToThird
	fmt.Println("Third")
	end <- 1
}

func main() {
	for i := 0; i < 100; i++ {
		go printFirst()
		go printSecond()
		go printThird()
		<- end
	}

	close(oneToTwo)
	close(twoToThird)
	close(end)
}
