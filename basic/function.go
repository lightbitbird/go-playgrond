package main

import (
    "fmt"
    "os"
)

func main() {
    x := 3
    y := 4
    z := 5

    // single output
    max_xy := max(x, y)
    fmt.Println("max=", max_xy)
    max_yz := max(y, z)
    fmt.Println("max=", max_yz)

    // multi outputs
    sum, multi := sumAndMultiply(x, y)
    fmt.Println("sum=", sum)
    fmt.Println("multyply=", multi)

    // multi outputs with names
    sum2, multi2 := multiWithOutputNames(y, z)
    fmt.Println("sum=", sum2)
    fmt.Println("multyply=", multi2)

    // multi arguments
    displaySumAll(2, 4, 7, 9, 11)

    // call by value
    fmt.Println("call by value ---------")
    fmt.Println("x=", x)
    x1 := add1(x) // just copy 'x' to arg
    fmt.Println("x + 1 =", x1)
    fmt.Println("x=", x) // still x = 3

    // call by reference (Pointer)
    fmt.Println("call by reference[Pointer] ---------")
    fmt.Println("x=", x)
    fmt.Println("pass x's address &x as the argument to addPointer func ------")
    x2 := addPointer(&x)
    fmt.Println("x + 1 =", x2) // x = 4
    fmt.Println("x=", x)       // x = 4
    fmt.Println()

    fmt.Println("defer ---------")
    for i := 0; i < 5; i++ {
        defer fmt.Printf("%d ", i)
    }
    write()
    read()
}

func add1(a int) int {
    a = a + 1 // change value 'a'
    return a  // return new value 'a'
}

// call by reference [Pointer]
// pointer variable *int
func addPointer(a *int) int {
    *a = *a + 1 // change value 'a'
    return *a   // return new value 'a'
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// multi output
func sumAndMultiply(A, B int) (int, int) {
    return A + B, A * B
}

// with output name
func multiWithOutputNames(A, B int) (add int, multiplied int) {
    add = A + B
    multiplied = A * B
    return
}

func displaySumAll(values ...int) {
    sum := 0
    for _, value := range values {
        fmt.Printf("arg %d in range\n", value)
        sum += value
    }
    fmt.Println("multi arguments sum=", sum)
}

func write() {
    file := "file.txt"
    fl, err := os.Create(file)
    if err != nil {
        fmt.Println(file, err)
    }
    defer fl.Close()
    for i := 0; i < 10; i++ {
        fl.WriteString("Just a test!\n")
        fl.Write([]byte("Just a test!\n"))
    }
}

func read() {
    file := "file.txt"
    fl, err := os.Open(file)
    if err != nil {
        fmt.Println(file, err)
    }
    defer fl.Close()
    buf := make([]byte, 1024)
    for {
        n, _ := fl.Read(buf)
        if 0 == n {
            break
        }
        os.Stdout.Write(buf[:n])
    }
}
