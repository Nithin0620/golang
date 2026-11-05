package main
// channels are used to communicate between goroutines and to synchronize them
//they have functionalities such as { Hold Data, Thread Safe, Listen for Data }

import "fmt"

func main () {
	var c = make(chan int)
	c <- 10 // this is how we can send data to a channel and this will block the goroutine until the data is received by another goroutine
	var i = <-c // this is how we can receive data from a channel and this will block the goroutine until the data is sent by another goroutine
	fmt.Println(i) // this will print 10 as we have sent 10 to the channel and received it in the variable i

	// this is gonna return a deadlock error as we are trying to send data to a channel but there is no goroutine to receive it and hence the goroutine is blocked and the program is deadlocked



	var c2 = make(chan int)
	go process(c2) 
	for i:= range c2{
		fmt.Println(i)// this is how we can receive data from a channel and this will block the goroutine until the data is sent by another goroutine and when the channel is closed it will exit the loop
	}

}

func process(c chan int){
	defer close(c) // this is another way how we can close a channel and it says like "do this stuff right before the channel exits" 
	for i:=0; i<5;i++ {
		c <- i // this is how we can send data to a channel and this will block the goroutine until the data is received by another goroutine
	}
	close(c) // this is how we can close a channel and this will unblock the goroutine that is waiting to receive data from the channel and it will exit the loop
	// if we dont add close then the proghram will be in deadlock again and it will go in infinite waiting waiting for more messafves 

}

var c = make(chan int,5) // this is how we can create a buffered channel and this will allow us to send data to the channel without blocking the goroutine until the buffer is full and when the buffer is full it will block the goroutine until the data is received by another goroutine
// this channel buffers can hold upto 5 values in the buffer and when we try to send the 6th value it will block the goroutine until the data is received by another goroutine