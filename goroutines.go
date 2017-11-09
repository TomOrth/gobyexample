package main

import "fmt"

func f(from string){
	for i := 0; i < 3; i++{
		fmt.Println(from, ":", i);
	}
}

func main(){
	f("direct")

	go f("goroutine") //goroutine example

	go func(msg string){//anonymous function
		fmt.Println(msg)
	}("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}