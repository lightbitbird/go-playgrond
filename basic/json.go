package main

import (
    "encoding/json"
    "fmt"
)

func main() {
    var s ServerSlice
    str := `{"servers":[{"serverName":"Shanghai_VPN", "serverIP":"27.0.0."},{"serverName":"Beijing_VPN","serverIP":"27.0.0.2"}]}`
    json.Unmarshal([]byte(str), &s)
    fmt.Println(s)
    fmt.Println()

    // from empty interface
    fmt.Println("Unmarshal json from empty interface{} -------------")
    b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
    var f interface{}
    err := json.Unmarshal(b, &f)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(f)
    }
    m := f.(map[string]interface{})
    for k, v := range m {
        switch vv := v.(type) {
        case string:
            fmt.Println(k, "is string", vv)
        case int:
            fmt.Println(k, "is int", vv)
        case float64:
            fmt.Println(k,"is float64",vv)
        case []interface{}:
            fmt.Println(k, "is an array:")
            for i, u := range vv {
                fmt.Println(i, u)
            }
        default:
            fmt.Println(k, "is of a type I don't know how to handle")
        }
    }
    fmt.Println()

    fmt.Println("Marshal struct to json -------------")
    s.Servers = append(s.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
    s.Servers = append(s.Servers, Server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
    b2, err2 := json.Marshal(s)
    if err2 != nil {
        fmt.Println("json err: ", err2)
    }
    fmt.Println(string(b2))
}

type Server struct {
    ServerName string `json:"serverName"`
    ServerIP string `json:"serverIP"`
}

type ServerSlice struct {
    Servers []Server `json:"servers"`
}
