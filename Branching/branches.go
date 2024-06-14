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

	//Deferred Functions: are called in LIFO order after main function completes
	//	Usually we use defer statements to release the resources
	//  we want to release them in opposite order in which we acquirred them(LIFO)
	fmt.Println("Main1")
	defer fmt.Println("Defer 1")
	fmt.Println("Main2")
	defer fmt.Println("Defer 2")
	// Main1
	// Main2
	// Defer 2
	// Defer 1

	//Panic and Recover
	fmt.Printf("Dividing 10 with 5: %v", divide(10, 5))
	fmt.Printf("Dividing 10 with 0: %v", divide(10, 0))
	//Before adding defer anonymous function in divide method
	//Dividing 10 with 5: 2Defer 2
	// Defer 1
	// panic: runtime error: integer divide by zero

	// goroutine 1 [running]:
	// main.divide(...)
	// 		C:/Users/ajpata/OneDrive - SAS/Training/GO/Branching/branches.go:39
	// main.main()
	// 		C:/Users/ajpata/OneDrive - SAS/Training/GO/Branching/branches.go:34 +0x22b
	// exit status 2
}

func divide(divisor, dividend int) int {
	defer func() {
		msg := recover()
		if msg != nil {
			fmt.Println(msg)
		}
	}()
	//After adding defer anonymous function
	// Dividing 10 with 5: 2runtime error: integer divide by zero
	// Dividing 10 with 0: 0Defer 2
	// Defer 1
	return divisor / dividend
}
