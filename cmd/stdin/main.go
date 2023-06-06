package main

import (
	"fmt"
	"go-utils/pkg/stdin"
)

func main() {
    // fmt.Print("Input X:")
    // t := stdin.StrStdin()
    // fmt.Println("Output X:", t)

    fmt.Print("Input Y:")
    t, err := stdin.IntStdin()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("output", t)
}
