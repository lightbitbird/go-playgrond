package main

import (
    "net/http"
    "fmt"
    "strings"
    "log"
    "html/template"
)

func main() {
    http.HandleFunc("/", sayHelloName)
    http.HandleFunc("/login", login)
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

func sayHelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm() // analyze option
    fmt.Println(r.Form)
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Hello world!")
}

func login(w http.ResponseWriter, r *http.Request) {
    r.ParseForm() // analyze option
    fmt.Println("method:", r.Method)
    fmt.Println("Form: ", r.Form)
    if r.Method == "GET" {
        t := template.Must(template.ParseFiles("web/login.gtpl"))
        t.Execute(w, nil)
    } else {
        m := map[string]string {
            // output to a client with html escape strings
            "username": template.HTMLEscaper(r.Form["username"]),
            "password": "*******",
        }

        t := template.Must(template.ParseFiles("web/login_result.gtpl"))
        // draw the template
        if err := t.ExecuteTemplate(w, "login_result.gtpl", m); err != nil {
            log.Fatal(err)
        }
        // output to a client with a raw html
        t2, _ := template.New("username").Parse(`{{define "T"}}{{.}}{{end}}`)
        t2.ExecuteTemplate(w, "T", template.HTML("<script>alert('Hello, go lang!')</script>"))
    }
}
