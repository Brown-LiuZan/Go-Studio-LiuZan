package main

import (
    "fmt"
    "os"
)

func main() {
    for i, arg := range os.Args[1:] {
        fmt.Println(i, ": ", arg)
    }

    //var s, sep string

    //for _, arg := range os.Args[1:] {
    //    s += sep + arg
    //    sep = " "
    //}
    //fmt.Println(s)

    //for i:= 1; i < len(os.Args); i++ {
    //    s += sep + os.Args[i]
    //    sep = " "
    //}
    //fmt.Println(s)

    //i := 1
    //for ; i < len(os.Args); i++ {
    //    s += sep + os.Args[i]
    //    sep = " "
    //}
    //fmt.Println(s)

    //i := 1
    //for i < len(os.Args) {
    //    s += sep + os.Args[i]
    //    sep = " "
    //    i++
    //}
    //fmt.Println(s)
}
