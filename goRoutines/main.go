package main

import (
	"fmt"
	"time"
)

func main() {
	// go routines and channels
	c := make(chan int)
	go doSomething(c)
	<-c

	// pointers
	g := 25
	fmt.Printf("g = %d\n", g)
	h := &g
	fmt.Printf("index memory &g or h: %v\n", h)
	fmt.Printf("*h: %d\n", *h)
}

func doSomething(c chan int) {
	time.Sleep(2 * time.Second)
	fmt.Println("Done")
	c <- 1
}
