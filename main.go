package main

import (
	"./addtwoint"
	"fmt"
)

func main() {
	fmt.Printf("Functions-in-Go\n\n")

    fmt.Println("Example: package name: addtwoint, function name: AddTwoInt")
    fmt.Printf("addtwoint.AddTwoInt(1, 2) is %d\n", addtwoint.AddTwoInt(1, 2))
    fmt.Println()
}
