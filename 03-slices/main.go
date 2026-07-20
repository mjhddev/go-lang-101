package main

import "fmt"

func main() {

	numbers := []int{10, 20, 30}

	numbers = append(numbers, 40)
	numbers = append(numbers, 50)

	fmt.Println("Numbers:", numbers)

	fmt.Println("\nUsing basic for loop")

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	fmt.Println("\nUsing range")

	for index, value := range numbers {
		fmt.Printf("Index: %d, value: %d\n", index, value)
	}

	fmt.Println("\nIgnoring index")

	for _, value := range numbers {
		fmt.Println(value)
	}
}
