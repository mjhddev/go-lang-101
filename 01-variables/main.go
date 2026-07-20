package main

import "fmt"

func main() {
	// Menggunakan keyword var
	var name string = "Mujahid"

	// Type inference
	var age = 25

	// Short variable declaration
	city := "Bandar Lampung"

	// Multiple variable declaration
	job, active := "Backend Developer", true

	fmt.Println("Name  :", name)
	fmt.Println("Age   :", age)
	fmt.Println("City  :", city)
	fmt.Println("Job   :", job)
	fmt.Println("Active:", active)
}
