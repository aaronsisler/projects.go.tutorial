package main

import "fmt"

func main() {
	var newColors map[string]string

	colors := map[string]string{
		"red":   "123456",
		"green": "789ABC",
	}

	fmt.Println(colors)
	fmt.Println(len(colors))

	printMap(colors)

	makeColors := make(map[string]string)

	fmt.Println(makeColors)
	fmt.Println("")

	makeColors["white"] = "taco"

	fmt.Println(makeColors)
	fmt.Println("")

	delete(makeColors, "white")

	fmt.Println(makeColors)
	fmt.Println("")

	fmt.Println(newColors)
	fmt.Println("")

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
