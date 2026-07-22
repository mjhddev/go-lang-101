package main

import "fmt"

func main() {
	colors1 := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	fmt.Println("Map Literal:")
	fmt.Println(colors1)

	colors2 := make(map[string]string)

	colors2["red"] = "#ff0000"
	colors2["green"] = "#00ff00"
	colors2["blue"] = "#0000ff"

	fmt.Println("\nMap with make():")
	fmt.Println(colors2)

	var colors3 map[string]string

	colors3 = make(map[string]string)

	colors3["red"] = "#ff0000"
	colors3["green"] = "#00ff00"
	colors3["blue"] = "#0000ff"

	fmt.Println("\nMap with var:")
	fmt.Println(colors3)
}
