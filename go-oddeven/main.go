package main

import "fmt"

func main() {

	numbers := []int{}
	for i := 0; i <= 10; i++ {
		numbers = append(numbers, i)
	}

	for _, number := range numbers {
		if number%2 == 0 {
			fmt.Println(number, "is even")
		} else {
			fmt.Println(number, "is odd")
		}
	}

	fmt.Println("alternative, shorter solution")

	nums := make([]int, 11)

	for i := range nums {
		if i%2 == 0 {
			fmt.Printf("%v is even\n", i)
			continue
		}
		fmt.Printf("%v is odd\n", i)
	}
}
