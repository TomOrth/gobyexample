package main

import "fmt"

func main() {
    i := 1
    //single condition
    for i <= 3 {
        fmt.Println(i)
        i++
    }

    //classic for loop
    for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }

    //no condition until break or return
    for {
        fmt.Println("loop")
        break
    }

    for n := 0; n <= 5; n++ {
        if n%2 == 0{
            continue
        }
        fmt.Println(n)
    }
}
