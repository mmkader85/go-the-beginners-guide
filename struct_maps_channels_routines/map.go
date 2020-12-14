package main

import "fmt"

func main() {
	colors := map[string]string{
		"red" : "#ff0000",
		"green" : "#00ff00",
		"blue" : "#0000ff",
	}

	addColorToMap(colors, "yellow", "#00ffff")
	deleteColorFromMap(colors, "green")
	printMap(colors)

}

func addColorToMap(m map[string]string, color string, hex string) {
	m[color] = hex
}

func deleteColorFromMap(m map[string]string, color string) {
	delete(m, color)
}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Printf("Hex code for %s is %s\n", color, hex)
	}
}