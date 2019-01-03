package main

import (
    "fmt"
    "errors"
    "math/rand"
    "time"
    "strconv"
)

func main() {
    // variables
    var name1, name2, name3 = 1, 2, 3
    fmt.Printf("name1 = %d, name2 = %d, name3 = %d\n", name1, name2, name3)

    vname1, vname2, vname3 := 4, 5, 6
    fmt.Printf("vname1 = %d, vname2 = %d, vname3 = %d\n", vname1, vname2, vname3)

    _, b := 7, 8
    fmt.Printf("b = %d\n", b)

    var (
        i      int
        pi     float32
        prefix string
    )
    i = 10;
    pi = 3.14;
    prefix = "prefix"
    fmt.Printf("%d, %v, %s\n", i, pi, prefix)

    // constant variable
    const constName = "const name"
    const Pi float32 = 3.1415926
    fmt.Printf("constName = %s, Pi = %v\n", constName, Pi)

    // complex variable
    var c complex64 = 3 + 6i
    fmt.Printf("Value is: %v", c)

    // string
    str1 := "World"
    fmt.Printf("Hello %s\n", str1)
    str2 :=
        `Hello, 
      World.`
    fmt.Println(str2)

    // How to change string values
    // Go cannot change string values
    // change a string value with doing following steps.
    before := "morning"
    chr := []byte(before) // Cast to byte
    chr[0] = 'c'
    after := string(chr) // Cast to string
    fmt.Printf("%s\n", after)

    // concat strings
    s := "hello,"
    m := " world"
    a := s + m
    fmt.Printf("%s\n", a)

    // error type
    err := errors.New("an error occured")
    if err != nil {
        fmt.Println(err)
    }

    // iota enum
    const (
        x = iota // x == 0
        y = iota // y == 1
        z = iota // z == 2
        w
    )
    const v = iota
    fmt.Printf("x=%v, y=%v, z=%v, w=%v, v=%v\n", x, y, z, w, v)

    // array
    // var arr [n]type
    // 配列の長さを変更はできない
    // 配列の代入は値渡しで、配列のコピーであり、ポインタではない
    var arr [10]int
    arr[0] = 42
    arr[1] = 24
    fmt.Println("array ----------------------")
    fmt.Printf("The first element is %d\n", arr[0])
    fmt.Printf("The second element is %d\n", arr[1])
    fmt.Printf("The last element is %d\n", arr[9])

    // another definitions
    a1 := [10]int{1, 2, 3}
    fmt.Printf("The last element is %d\n", a1[0])
    fmt.Printf("The second element is %d\n", a1[1])
    fmt.Printf("The last element is %d\n", a1[9])

    // omission of an array size
    b1 := [...]int{4, 5, 6}
    for cnt := 0; cnt < len(b1); cnt++ {
        fmt.Printf("The element is %d\n", b1[cnt])
    }

    // 2 dimensions arrays
    // doubleArray := [2][3]int{ [3]int{1, 2, 3}, [3]int{5, 6, 7} }
    doubleArray := [2][3]int{{1, 2, 3}, {5, 6, 7}}
    for cnt := 0; cnt < len(doubleArray); cnt++ {
        for cnt2 := 0; cnt2 < len(doubleArray[cnt]); cnt2++ {
            fmt.Printf("[%d][%d]: %d\n", cnt, cnt2, doubleArray[cnt][cnt2])
        }
    }

    // if
    nine := 9
    if nine > 10 {
        fmt.Println("nine is greater than 10")
    } else {
        fmt.Println("nine is less than 10")
    }

    // define variable 'x' in only a condition scope
    if x := randomNumber() + 10; x > 10 {
        fmt.Println("x is greature than 10.", x)
    } else {
        fmt.Println("x is less than 10.", x)
    }

    // goto
    gotoFunc()

    // for
    fmt.Println("for --------")
    sm := 0
    for i := 0; i < 10; i++ {
        sm += i
    }
    fmt.Println("sum is equal to ", sm)

    // for : omission version
    sm2 := 1
    for ; sm2 < 40; {
        sm2 += sm2
    }
    fmt.Println("sum is equal to ", sm2)
    fmt.Println()

    fmt.Println("for range --------")
    kv := make(map[string]string)
    for i := 0; i < 10; i++ {
        index := strconv.Itoa(i)
        kv["k"+index] = "v" + index
    }
    for k, v := range kv {
        fmt.Printf("%s=%s\n", k, v)
    }

}

func randomNumber() int {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(10)
}

// goto loop func
func gotoFunc() {
    a := 0
    i := 1
LOOP:
    if i > 100 {
        goto BREAK
    }
    a += i
    i += 1
    goto LOOP
BREAK:
    fmt.Println(a)
}
