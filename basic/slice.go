package main

import (
    "fmt"
)

func main() {
    var ar = [6]byte{'a', 'b', 'c', 'd', 'e', 'f'}
    var slice1, slice2, slice3, slice4 []byte
    // all
    slice1 = ar[:]
    fmt.Printf("slice1=%v\n", slice1)

    // index 3~5
    slice2 = ar[3:5]
    fmt.Printf("slice2=%v\n", slice2)

    // index 3~
    slice3 = ar[3:]
    fmt.Printf("slice3=%v\n", slice3)

    // index ~3
    slice4 = ar[:3]
    fmt.Printf("slice4=%v\n", slice4)

    num := [...]int{1, 5, 4, 2, 3}
    slice5 := num[1:4]
    fmt.Println("slice5=", slice5)
    fmt.Println("len=", len(slice5))

    // slice is [Call by reference]
    plusOne(num[:])
    fmt.Println("call by reference: ", num)

    // capacity
    capnum := [...]int{1, 2, 3, 4, 5}
    capslice := capnum[1:4]
    fmt.Println("capslice=", capslice)
    fmt.Println("len=", len(capslice))
    fmt.Println("cap(capslice)=", cap(capslice))

    fmt.Println("slice[1:4] from capslice ------")
    capslice2 := capslice[1:4]
    fmt.Println("capslice2=", capslice2)
    fmt.Println("len=", len(capslice2))
    fmt.Println("cap=", cap(capslice2))

    fmt.Println("append ------")
    s1 := []int{1, 2, 3, 4, 5}
    fmt.Println("s1=", s1)

    s2 := append(s1, 6, 7)
    fmt.Println("s2(append 6, 7)=", s2)

    s3 := append(s1, s2...)
    fmt.Println("s3(append s2)=", s3)

    fmt.Println("copy ------")
    src1 := []int{1, 2}
    dest := []int{97, 98, 99}
    count := copy(dest, src1)
    fmt.Println("copy count(copied elements count)=", count)
    fmt.Println("dest=", dest)
    fmt.Println()

    src2 := []int{3}
    count = copy(dest[2:], src2)
    fmt.Println("copy count(copied elements count)=", count)
    fmt.Println(dest)

    fmt.Println("make ------")
    mk := make([]string, 5, 10)
    fmt.Println("mk=", mk)
    fmt.Println("len=", len(mk))
    fmt.Println("cap=", cap(mk))
    fmt.Println()

    mk2 := make([]string, 5)
    fmt.Println("mk2=", mk2)
    mk2[0] = "one"
    mk2[1] = "two"
    mk2[2] = "three"
    mk2[3] = "four"
    mk2[4] = "five"
    fmt.Println("mk2=", mk2)
    fmt.Println("len=", len(mk2))
    fmt.Println("cap=", cap(mk2))
    fmt.Println()

    fmt.Println("for range --------")
    for n, v := range mk2 {
        fmt.Printf("%d=%s\n", n, v)
    }
}

func plusOne(vals []int) {
    for i := 0; i < len(vals); i++ {
        vals[i] += 1
    }
}
