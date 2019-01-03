package main

import (
    "math"
    "fmt"
)

func main() {
    // method: func (c ReceiverType) funcName(parameters) (results)
    fmt.Println("method: func (c ReceiverType) funcName(parameters) (results)")
    r1 := Rectangle{12, 2}
    r2 := Rectangle{9, 4}
    c1 := Circle{10}
    c2 := Circle{25}

    fmt.Println("Area of r1 is: ", r1.area())
    fmt.Println("Area of r2 is: ", r2.area())
    fmt.Println("Area of c1 is: ", c1.area())
    fmt.Println("Aread of c2 is: ", c2.area())
    fmt.Println()

    fmt.Println("extends method ----------")
    mark := Student2{Human2{"Mark", 25, "222-2222-YYYY"}, "MIT"}
    sam := Employee2{Human2{"Sam", 45, "5555-7777-XXXX"}, "Golang Inc"}
    mark.SayHi()
    sam.SayHi()
}

type Rectangle struct {
    width, height float64
}

type Circle struct {
    radius float64
}

func (r Rectangle) area() float64 {
    return r.width * r.height
}

func (c Circle) area() float64 {
    return c.radius * c.radius * math.Pi
}

type Human2 struct {
    name  string
    age   int
    phone string
}

type Student2 struct {
    Human2 // anonymous field
    school string
}

type Employee2 struct {
    Human2 // anonymous field
    company string
}

func (h *Human2) SayHi() {
    fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}
