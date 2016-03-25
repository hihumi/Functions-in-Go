package main

import (
	"./addtwoint"
	"./subtwoint"
	"fmt"
)

func main() {
	fmt.Printf("Functions-in-Go\n\n")

	// addtwoint
	fmt.Println("Example: package name: addtwoint, function name: AddTwoInt")
	fmt.Printf("addtwoint.AddTwoInt(1, 2) is %d\n", addtwoint.AddTwoInt(1, 2))

	// Assignment
	myAddTwoInt := addtwoint.AddTwoInt(100, 200)
	fmt.Printf("Variable name: myAddTwoInt is ")
	fmt.Println(myAddTwoInt)
	fmt.Println()

	// subtwoint
	fmt.Println("Example: package name: subtwoint, function name: SubTwoInt")
	fmt.Printf("subtwoint.SubTwoInt(2, 1) is %d\n", subtwoint.SubTwoInt(2, 1))

	// Assignment
	mySubTwoInt := subtwoint.SubTwoInt(200, 100)
	fmt.Printf("Variable name: mySubTwoInt is ")
	fmt.Println(mySubTwoInt)
	fmt.Println()
}
