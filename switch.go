package main

import "fmt"
import "time"

func main() {
    i := 2
    fmt.Println("Write ", i, " as ");
    switch i {
        case 1:
            fmt.Println("one")
        case 2:
            fmt.Println("two")
        case 3:
            fmt.Println("three")
    }

    switch time.Now().Weekday() {
        case time.Saturday, time.Sunday:
            fmt.Println("weekend")
        default:
            fmt.Println("hell")
    }

    t := time.Now()
    switch {
        case t.Hour() < 12:
            fmt.Println("Its before noon")
        default:
            fmt.Println("afternoon m8")
    }

    whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
}
