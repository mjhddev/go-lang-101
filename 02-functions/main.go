package main

import "fmt"

func greet(name string) {
	fmt.Println("Hello,", name)
}

func add(a, b int) int {
	return a + b
}

func getUser() (string, int) {
	return "Mujahid", 25
}

func main() {
	greet("Mujahid")
	result := add(10, 5)
	fmt.Println("Result:", result)

	name, age := getUser()
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
}
