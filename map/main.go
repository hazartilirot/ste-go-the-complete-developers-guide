package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
		"black": "#000000",
	}

	fmt.Println(colors)

	var moreColors map[string]string     // nil
	moreColors = make(map[string]string) // initializing
	moreColors["white"] = "#ffffff"

	fmt.Println(moreColors)

	lotsOfColors := make(map[string]string)

	lotsOfColors["black"] = "#000000"

	fmt.Println(lotsOfColors)

	delete(colors, "red")

	fmt.Println(colors)

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
