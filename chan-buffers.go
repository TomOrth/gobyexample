package main

import "fmt"

func main(){
	messages := make(chan string, 2) //limited input

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}