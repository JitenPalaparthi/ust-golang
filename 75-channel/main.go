package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		ch <- 100 // sender sending data <- towards the channel
	}()
	v := <-ch // receiver, it is from the channel
	fmt.Println(v)
}

// 1. channel is kind of coundit
// 2. there are buffered and unbuffered channels
// 3. there is a sender and a receiver
// 4. channel can be nil
// 5. make is built-in function we use for channels
// 6. ch:= make(chan int), this is unbuffered channel, this channel can take one integer value at a time
// 7. sending data --> ===== --> receiving
// 8. sender is blocked until the receiver receives the data
// 9. receiver is blocked until the sender sends the data
