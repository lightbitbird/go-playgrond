package main

import (
    "runtime"
    "fmt"
    "time"
)

func main() {
    // goroutine
    fmt.Println("goroutine -----------")
    go say("world")
    say("hello")
    fmt.Println()

    fmt.Println("channel -----------")
    fmt.Println("chan type: definition channel")
    fmt.Println("chan <- type: send specialized channel")
    fmt.Println("<- type: receive specialized channel")

    a := []int{7, 2, 8, -9, 4, 0}
    c := make(chan int)
    go sum(a[:len(a)/2], c)
    go sum(a[len(a)/2:], c)
    x, y := <-c, <-c // receive from c
    fmt.Println(x, y, x+y)
    fmt.Println()

    fmt.Println("beffered channels -------")
    ch := make(chan int, 1)     // 1 buffering
    time.After(5 * time.Second) // timeout set
    ch <- 1
    fmt.Println(<-ch)
    ch <- 2 // block buffering, so error will occur without timeout setting
    fmt.Println(<-ch)
    fmt.Println()

    fmt.Println("select -------")
    sch := make(chan int)
    quit := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(<-sch)
        }
        quit <- 0
    }()
    fibonacci(sch, quit)
    fmt.Println()

    fmt.Println("data sharing between goroutine -------")
    // data sharing channel
    counter := make(chan int)
    // goroutine for end notice channel
    end := make(chan bool)
    for i := 0; i < goroutines; i++ {
        go func(counter chan int) {
            val := <-counter
            val++
            fmt.Println("counter = ", val)
            if val == goroutines {
                end <- true
            }
            counter <- val
        }(counter)
    }
    counter <- 0
    <-end
    fmt.Println("end")
}

const goroutines = 5

func say(s string) {
    for i := 0; i < 5; i++ {
        runtime.Gosched()
        fmt.Println(s)
    }
}

func sum(a []int, c chan int) {
    total := 0
    for _, v := range a {
        total += v
    }
    c <- total // send total to c
}

func fibonacci(c, quit chan int) {
    x, y := 1, 1
    for {
        select {
        case c <- x:
            x, y = y, x+y
        case <-quit:
            fmt.Println("quit")
            return
        }
    }
}
