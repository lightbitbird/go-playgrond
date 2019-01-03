package main

import (
    "fmt"
    "os"
)

func main() {
    var readFunc ReadFunc
    var dict Dictionary
    // 同じ戻り値、引数でtypeと関数を紐付けることができる
    readFunc = readOut
    dict.name = "Glaetzer Amon-ra Shiraz"
    dict.place = "Australia"
    fmt.Println(readFunc(dict))
    fmt.Println()

    fmt.Println("type as a same interface -------")
    slice := []int{1, 2, 3, 4, 5, 6}
    fmt.Println("slice = ", slice)
    odd := filter(slice, isOdd) // call by value
    fmt.Println("Odd elements of slice are ", odd)
    even := filter(slice, isEven) // call by value
    fmt.Println("Even elements of slice are: ", even)
    fmt.Println()

    fmt.Println("Panic & Recover -------")
    // fmt.Println("throws panic = ", throwsPanic(fpanic))
    fmt.Println("Recover from panic = ", recoverFromPanic(fpanic))
}

var env = os.Getenv("NO_ENV")

type Dictionary struct {
    name  string
    place string
}

type ReadFunc func(Dictionary) string

func readOut(d Dictionary) string {
    return fmt.Sprintf("[%s] is made in [%s].", d.name, d.place)
}

type testInt func(int) bool // definition of type func

func isOdd(integer int) bool {
    if integer%2 == 0 {
        return false
    }
    return true
}

func isEven(integer int) bool {
    if integer%2 == 0 {
        return true
    }
    return false
}

func filter(slice []int, f testInt) []int {
    var result []int
    for _, value := range slice {
        if f(value) {
            result = append(result, value)
        }
    }
    return result
}

func fpanic() {
    fmt.Println("env=", env)
    if env == "" {
        panic("no value for $NO_ENV")
    }
}

func throwsPanic(f func()) (b bool) {
    f() // panicを出現させる
    return
}

func recoverFromPanic(f func()) (b bool) {
    defer func() {
        // panicが出現すると復元を行う
        if x := recover(); x != nil {
            b = true
        }
    }()
    f() // panicを出現させる
    return
}
