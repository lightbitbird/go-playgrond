package main

import (
    "os"
    "fmt"
)

func main() {
    // directory
    fmt.Println("create directory -----------")
    os.Mkdir("dir1", 0777)
    os.MkdirAll("dir1/dir2/test", 0777)
    err := os.Remove("dir1")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("remove directories -----------")
    os.RemoveAll("dir1")
    fmt.Println()

    fmt.Println("create a file, write -----------")
    useFile := "test10.txt"
    fout, err := os.Create(useFile)
    if err != nil {
        fmt.Println(useFile, err)
        return
    }
    defer fout.Close()
    for i := 0; i < 10; i++ {
        fout.WriteString("Just a test!\n")
        fout.Write([]byte("Just a test!\n"))
    }
    fmt.Println()
    fmt.Println("open a file, read -----------")
    fl, err1 := os.Open(useFile)
    if err1 != nil {
        fmt.Println(useFile, err1)
        return
    }
    defer fl.Close()
    buf := make([]byte, 1024)
    for {
        n, _ := fl.Read(buf)
        if 0 == n {
            break
        }
        os.Stdout.Write(buf[:n])
    }
    fmt.Println("remove a file -----------")
    os.Remove("test10.txt")
}
