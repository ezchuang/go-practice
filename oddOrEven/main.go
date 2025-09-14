package main

import "fmt"

func main() {
	ints := []int{}

	for i := int(0); i < 11; i++ {
		ints = append(ints, i)
	}

	for _, v := range ints {
		if v%2 == 0 {
			fmt.Println(v, "is even")
		} else {
			fmt.Println(v, "is odd")
		}
	}
}
