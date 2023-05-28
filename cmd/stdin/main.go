package main

import (
	"fmt"
	"go-utils/pkg/stdin"
)

func main() {
    fmt.Print("Input X:")
    t := stdin.StrStdin()
    fmt.Println("Output X:", t)
}
