package main

import "fmt"

func main() {
    var tom person
    // initialize 1
    tom.name, tom.age = "Tom", 18

    // initialize 2
    bob := person{age: 25, name: "Bob"}

    // initilize 3
    paul := person{"Paul", 43}

    tb, tb_diff := older(tom, bob)
    tp, tp_diff := older(tom, paul)
    bp, bp_diff := older(bob, paul)

    fmt.Printf("Of %s and %s, %s is older by %d years\n", tom.name, bob.name, tb.name, tb_diff)
    fmt.Printf("Of %s and %s, %s is older by %d years\n", tom.name, paul.name, tp.name, tp_diff)
    fmt.Printf("Of %s and %s, %s is older by %d years\n", bob.name, paul.name, bp.name, bp_diff)
    fmt.Println()

    fmt.Println("anonymous field inside struct -------")
    // initialize a student jane
    jane := Student{Human: Human{name: "Jane", age: 35, weight: 100}, speciality: "Biology"}
    fmt.Println("Her name is ", jane.name)
    fmt.Println("Her age is ", jane.age)
    fmt.Println("Her weight is ", jane.weight)
    fmt.Println("Her speciality is ", jane.speciality)

    // Add jane's skills
    jane.Skills = []string{"anatomy", "math"}
    fmt.Println("Her skills are ", jane.Skills)
    fmt.Println("She acquired two new ones ")
    jane.Skills = append(jane.Skills, "physics", "golang")
    fmt.Println("Her skills now are ", jane.Skills)

    // Add a value to anonymous field int
    jane.int = 3
    fmt.Println("Her preferred number is ", jane.int)
    fmt.Println()

    fmt.Println("Same field name between parent and child struct -------")
    Max := Employee{Human{"Bob", 34, 65, "000-1111-XXXX"}, "Designer", "333-2222-ZZZZ"}
    fmt.Println("Max's work phone is: ", Max.phone)
    fmt.Println("Max's personal phone is: ", Max.Human.phone)
}

type person struct {
    name string
    age  int
}

func older(p1, p2 person) (person, int) {
    if p1.age > p2.age {
        return p1, p1.age - p2.age
    }
    return p2, p2.age - p1.age
}

type Skills []string

type Human struct {
    name   string
    age    int
    weight int
    phone  string
}

type Student struct {
    Human  // anonymous field, struct
    Skills // anonymous field, string slice
    int    // anonymous field, int
    speciality string
}

type Employee struct {
    Human // anonymous field, struct
    speciality string
    phone      string // staff phone filed
}
