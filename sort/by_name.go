package main

import (
    "log"
    "sort"
)

type Planet struct {
    Name       string  `json:"name"`
    Aphelion   float64 `json:"aphelion"`   // in million km
    Perihelion float64 `json:"perihelion"` // in million km
    Axis       int64   `json:"Axis"`       // in km
    Radius     float64 `json:"radius"`
}

type By func(p1, p2 *Planet) bool

func (by By) Sort(planets []Planet) {
    ps := &planetSorter{
        planets: planets,
        by:      by, // The Sort method's receiver is the function (closure) that defines the sort order.
    }
    sort.Sort(ps)
}

type planetSorter struct {
    planets []Planet
    by      func(p1, p2 *Planet) bool // Closure used in the Less method.
}

// Len is part of sort.Interface.
func (s *planetSorter) Len() int {
    return len(s.planets)
}

// Swap is part of sort.Interface.
func (s *planetSorter) Swap(i, j int) {
    s.planets[i], s.planets[j] = s.planets[j], s.planets[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *planetSorter) Less(i, j int) bool {
    return s.by(&s.planets[i], &s.planets[j])
}

func main() {
    var mars = new(Planet)
    mars.Name = "Mars"
    mars.Aphelion = 249.2
    mars.Perihelion = 206.7
    mars.Axis = 227939100
    mars.Radius = 3389.5

    var earth = new(Planet)
    earth.Name = "Earth"
    earth.Aphelion = 151.930
    earth.Perihelion = 147.095
    earth.Axis = 149598261
    earth.Radius = 6371.0

    var venus = new(Planet)
    venus.Name = "Venus"
    venus.Aphelion = 108.939
    venus.Perihelion = 107.477
    venus.Axis = 108208000
    venus.Radius = 6051.8

    planets := []Planet{*mars, *venus, *earth}

    By(func(p1, p2 *Planet) bool {
        return p1.Name < p2.Name
    }).Sort(planets)

    log.Println(planets)

    By(func(p1, p2 *Planet) bool {
        return p1.Axis < p2.Axis
    }).Sort(planets)

    log.Println(planets)

    // sort.Slice
    sort.Slice(planets, func(i, j int) bool {
        return planets[i].Name < planets[j].Name
    })

    log.Println(planets)
}
