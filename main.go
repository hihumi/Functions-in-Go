package main

import (
	"./addtwoint"
	"./divtwoint"
	"./evenintlimit100"
	"./evenintoroddint"
	"./multitwoint"
	"./subtwoint"
	"fmt"
)

func main() {
	fmt.Printf("Functions-in-Go\n\n")

	// package: addtwoint
	fmt.Println("Example: package name: addtwoint, function name: AddTwoInt")
	fmt.Printf("addtwoint.AddTwoInt(1, 2) is %d\n", addtwoint.AddTwoInt(1, 2))
	// Assignment
	myAddTwoInt := addtwoint.AddTwoInt(100, 200)
	fmt.Printf("Variable name: myAddTwoInt is ")
	fmt.Println(myAddTwoInt)
	fmt.Println()

	// package: subtwoint
	fmt.Println("Example: package name: subtwoint, function name: SubTwoInt")
	fmt.Printf("subtwoint.SubTwoInt(2, 1) is %d\n", subtwoint.SubTwoInt(2, 1))
	// Assignment
	mySubTwoInt := subtwoint.SubTwoInt(200, 100)
	fmt.Printf("Variable name: mySubTwoInt is ")
	fmt.Println(mySubTwoInt)
	fmt.Println()

	// package: multitwoint
	fmt.Println("Example: package name: multitwoint, function name: MultiTwoInt")
	fmt.Printf("multitwoint.MultiTwoInt(3, 4) is %d\n", multitwoint.MultiTwoInt(3, 4))
	// Assignment
	myMultiTwoInt := multitwoint.MultiTwoInt(300, -400)
	fmt.Printf("Variable name: mySubTwoInt is ")
	fmt.Println(myMultiTwoInt)
	fmt.Println()

	// package: divtwoint
	fmt.Println("Example: package name: divtwoint, function name: DivTwoInt")
	fmt.Printf("divtwoint.DivtiTwoInt(3, 4) is %d\n", divtwoint.DivTwoInt(10, 5))
	// Assignment
	myDivTwoInt := divtwoint.DivTwoInt(100, -50)
	fmt.Printf("Variable name: myDivTwoInt is ")
	fmt.Println(myDivTwoInt)
	fmt.Println()

	// package: evenintoroddint
	fmt.Println("Example: package name: evenintoroddint, function name: EvenIntOrOddInt")
	evenintoroddint.EvenIntOrOddInt(10)
	evenintoroddint.EvenIntOrOddInt(11)
	evenintoroddint.EvenIntOrOddInt(-10)
	evenintoroddint.EvenIntOrOddInt(-11)
	fmt.Println()

	// package: evenintlimit100
	fmt.Println("Example: package name: evenintlimit100, function name: EvenIntLimit100")
	evenintlimit100.EvenIntLimit100(10) // This case start integer is 10.
	fmt.Println()
	evenintlimit100.EvenIntLimit100(-50) // This case start integer is -50.
	fmt.Println()
	evenintlimit100.EvenIntLimit100(100)
	fmt.Println()
	evenintlimit100.EvenIntLimit100(500)
}
