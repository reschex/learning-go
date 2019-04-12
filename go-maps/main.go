package main

import "fmt"

func main() {
	colours := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
	}

	// var colours map[string]string

	// colours := make(map[string]string)

	colours["white"] = "#ffffff"

	//delete(colours, "white")

	// fmt.Println(colours)

	printMap(colours)
}

func printMap(c map[string]string) {
	for colour, hex := range c {
		fmt.Printf("colour: %v, hex: %v\n", colour, hex)
	}
}
