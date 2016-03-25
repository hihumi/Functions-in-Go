package evenintoroddint

import "fmt"

func EvenIntOrOddInt(i int) int {
	for {
		if i%2 == 0 {
			fmt.Println(i, "is even")
			break
		} else {
			fmt.Println(i, "is odd")
			break
		}
	}

	return i
}
