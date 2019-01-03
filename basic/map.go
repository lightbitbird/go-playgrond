package main

import (
    "fmt"
    "strconv"
)

func main() {
    currencies := make(map[string]string)
    currencies["jpn"] = "JPY"
    currencies["usa"] = "USD"
    currencies["eu"] = "EUR"
    currencies["china"] = "CNY"

    fmt.Println("jpy=", currencies["jpn"])
    fmt.Println("china", currencies["CNY"])

    fmt.Println("all ------------")
    for country, currency := range currencies {
        fmt.Println(country, currency)
    }
    fmt.Println("----------------")
    fmt.Println()

    fmt.Println("exist ------------")
    currency, exist := currencies["rss"]
    if exist {
        fmt.Println("registered", currency)
    } else {
        fmt.Println("not yet registered")
    }
    fmt.Println()

    fmt.Println("add 'canada' ------------")
    currencies["canada"] = "USD/CAD"
    fmt.Println("currencies after add=", currencies)
    fmt.Println()

    fmt.Println("delete 'usa' ------------")
    delete(currencies, "usa")
    fmt.Println("currencies after delete=", currencies)
    fmt.Println()

    // initialize
    initMap := map[string]string{
        "jpn":   "JPY",
        "usa":   "USD",
        "eu":    "EUR",
        "china": "CNY",
    }
    fmt.Println(initMap)

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
