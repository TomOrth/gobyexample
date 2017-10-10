package main

import "fmt"

func fact(i, n int) int {
	if i == 0 {
		return n
	}
	return fact(i-1, n*i)
}

func main(){
	fmt.Println(fact(7,1))
}