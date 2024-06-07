package main

import "fmt"

func main() {
	i := 999
	if i == 999 {
		fmt.Println("I is 999")
	}
	
	switch i {
	case 1:
		fmt.Println("first case")
	case 2 + 3, 2 + i:
		fmt.Println("second case")
	default:
		fmt.Println("default case")
	}

}
