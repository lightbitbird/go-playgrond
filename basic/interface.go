package main

import (
    "fmt"
    "strconv"
    "reflect"
)

func main() {
    // interface on Golang is a combination of methods
    fmt.Println("interface on Golang is a combination of methods --------")

    mike := Student3{Human3{"Mike", 25, "222-2222-XXXX"}, "MIT", 0.00}
    paul := Student3{Human3{"Paul", 26, "111-1111-XXXX"}, "Harvard", 100}
    sam := Employee3{Human3{"Sam", 36, "333-4444-XXXX"}, "Golang Inc.", 1000}
    tom := Employee3{Human3{"Tom", 37, "444-5555-XXXX"}, "Things Ltd.", 5000}

    // type Men variable
    var i Men
    // store Student mike
    i = mike
    fmt.Println("This is Mike, Student:")
    i.SayHi()
    i.Sing("November rain")
    // result of String()
    fmt.Println("implements String() -- This human is : ", i)
    fmt.Println()

    // store Employee mike
    i = tom
    fmt.Println("This is Tom, an Employee:")
    i.SayHi()
    i.Sing("Born to be wild")
    // result of String()
    fmt.Println("implements String() -- This human is : ", i)
    fmt.Println()

    // slice for Men
    fmt.Println("make slice of Men and use them ---------")
    x := make([]Men, 4)
    x[0], x[1], x[2], x[3] = paul, sam, mike, tom

    for i, value := range x {
        fmt.Println("-----------------------")
        value.SayHi()
        index := strconv.Itoa(i)
        value.Sing("lyrics_" + index)
    }
    fmt.Println()

    // empty interface
    fmt.Println("empty interface ---------------------")
    var a interface{}
    it := 5
    s := "Hello world"
    a = it
    fmt.Println("empty interface 'a' = int : ", a)
    a = s
    fmt.Println("empty interface 'a' = string : ", a)
    fmt.Println()

    // empty interface
    fmt.Println("Comma-ok assertion ---------------------")
    list := make(List, 3)
    list[0] = 1
    list[1] = "Hello"
    list[2] = Person{"Dennis", 40}
    fmt.Println("if version -------------")
    for index, element := range list {
        if value, ok := element.(int); ok {
            fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
        } else if value, ok := element.(string); ok {
            fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
        } else if value, ok := element.(Person); ok {
            fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
        } else {
            fmt.Printf("list[%d] is of a different type\n", index)
        }
    }
    fmt.Println("switch version -------------")
    for index, element := range list {
        switch value := element.(type) {
        case int:
            fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
        case string:
            fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
        case Person:
            fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
        default:
            fmt.Printf("list[%d] is of a different type\n", index)
        }
    }
    fmt.Println()

    // reflection
    fmt.Println("reflection ---------------------")
    types := reflect.TypeOf(x)
    fmt.Println("reflect.TypeOf: ", types)
    vals := reflect.ValueOf(x)
    fmt.Println("reflect.ValueOf: ", vals)
    t := reflect.TypeOf(x[0])
    ft1 := t.Field(0).Type
    fmt.Println("field1 type: ", ft1)
    ft2 := t.Field(1).Type
    fmt.Println("field2 type: ", ft2)
    ft3 := t.Field(2).Type
    fmt.Println("field3 type: ", ft3)
    v := reflect.ValueOf(x[0])
    fv1 := v.Field(0)
    fmt.Println("field1 value: ", fv1)
    fv2 := v.Field(1)
    fmt.Println("field2 value: ", fv2)
    fv3 := v.Field(2)
    fmt.Println("field3 value: ", fv3)
    fmt.Println()

    xx := 3.4
    vv := reflect.ValueOf(xx)
    fmt.Println("type: ", vv.Type())
    fmt.Println("kind is float64: ", vv.Kind() == reflect.Float64)
    fmt.Println("value: ", vv.Float())
    // 代入するときはPointerで参照渡しにしなければ、reflectionでアドレスが見つけられない
    fmt.Println("reflect.ValueOf(&xx) ---------")
    p := reflect.ValueOf(&xx)
    pvv := p.Elem()
    pvv.SetFloat(7.1)
    fmt.Println("after changed value: ", pvv.Float())

}

// Interface Men is implemented by Human3, Student3, Employee3
// These three types are implementing the two methods
type Men interface {
    SayHi()
    Sing(lyrics string)
}

type Human3 struct {
    name  string
    age   int
    phone string
}

type Student3 struct {
    Human3 // anonymous field
    school string
    loan   float32
}

type Employee3 struct {
    Human3 // anonymous field
    company string
    money   float32
}

// implements SayHi mehod on Human3
func (h Human3) SayHi() {
    fmt.Printf("Hi, I'm %s you can call me on %s\n", h.name, h.phone)
}

// implements Sing mehod on Human3
func (h Human3) Sing(lyrics string) {
    fmt.Println("La la la la...", lyrics)
}

// overroad SayHi method on Employee
func (e Employee3) SayHi() {
    fmt.Printf("Hi, I'm %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone)
}

// implements Stringer String()
func (h Human3) String() string {
    return "❰" + h.name + " - " + strconv.Itoa(h.age) + " years -  ✆ " + h.phone + "❱"
}

type Element interface{}
type List []Element

type Person struct {
    name string
    age  int
}

func (p Person) String() string {
    return "name: " + p.name + " - age: " + strconv.Itoa(p.age) + " years)"
}
