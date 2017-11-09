package main

import "fmt"

func ping(pings chan <- string, msg string){ //channel can only send values
	pings <- msg
}

func pong(pings <-chan string, pongs chan <- string){ //ping is example of only getting values
	msg := <-pings
	pongs <- msg
}

func main(){
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}