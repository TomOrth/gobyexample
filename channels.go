package main

import "fmt"

func main(){
	messages := make(chan string) //generate a channel

	go func(){messages <- "ping"}() //feed into the channel

	msg  := <-messages //retrieve channel data
	fmt.Println(msg)
}