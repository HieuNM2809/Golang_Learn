package main

import "fmt"

func main() {
    a := 10
    b := 20
    fmt.Println(findMax(a, b))

    c := 30
    d := 40
    fmt.Println(findMax(c, d))
}

func findMax(a, b int) int {
    if a > b {
        return a
    }
    return b
}
