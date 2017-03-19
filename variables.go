package main

import "fmt"

func main() {
    //type comes after variable name
    var a string = "initial"
    fmt.Println(a)
    //multiple variable declaration
    var b, c int = 1,2
    fmt.Println(b, c)
    //type can be implied
    var d = true
    fmt.Println(d)
    //variables don't need an initial value (zero-valued)
    var e int
    fmt.Println(e)
    //:= is a shorthand for variable
    f := "short"
    fmt.Println(f)
}
