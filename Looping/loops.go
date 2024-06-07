package main

import (
	"fmt"
)

func main() {
	i := 0
	//Infinite loop
	for {
		fmt.Println(i)
		i += 1
		break
	} //Without break this loop runs till int is out of memory,exit

	//Conditional loop
	for i <= 10 {
		fmt.Println(i)
		i += 2
	}

	//Counter based loop
	for i = 1; i < 11; i++ {
		fmt.Println(2 * i)
	}

	//Looping with collections
	arr := [3]int{101, 102, 103}
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
