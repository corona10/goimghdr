package main

import (
        "fmt"
        "os"

        "github.com/corona10/goimghdr"
)

func main() {
        path1 := "../imgData/gopher.jpeg"
        path2 := "../imgData/gopher.png"
        ret1, _ := goimghdr.What(path1)
        fmt.Printf("%s: %s\n", path1, ret1)
        f, _ := os.Open(path2)
        ret2, _ := goimghdr.WhatFromReader(f)
        fmt.Printf("%s: %s\n", path2, ret2)
        f.Close()
}
