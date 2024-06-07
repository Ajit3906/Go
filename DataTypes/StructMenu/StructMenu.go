package main

import (
	"fmt"
	"strings"
)

func main() {
	var choice int
	fmt.Println("Enter the menu number-\n1.Menu")
	fmt.Scan(&choice)

	type Menu struct {
		name  string
		price map[string]float32
	}

	items := []Menu{{name: "coffee", price: map[string]float32{"small": 4.5, "big": 10}}, {name: "Expresso", price: map[string]float32{"single": 10.5, "double": 22.3}}}
	//fmt.Println(items) //[{coffee map[big:10 small:4.5]} {Expresso map[double:22.3 single:10.5]}]

	if choice == 1 {
		for _, item := range items {
			fmt.Println(item.name)
			fmt.Println(strings.Repeat("-", 20))
			for size, price := range item.price {
				fmt.Printf("\t%10s%10.2f\n", size, price)
			}
		}
	}
}
