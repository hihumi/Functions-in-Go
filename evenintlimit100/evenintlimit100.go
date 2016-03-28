package evenintlimit100

import "fmt"

func EvenIntLimit100(i int) int {
	fmt.Println(i)
	for {
		i++
		if i >= 100 {
			fmt.Println("limit 100")
			break
		}
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}

	return i
}
