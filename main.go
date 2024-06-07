package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//Variables and Data types
	a, b := 4, 2
	println("Hello world!!")
	c := a + b
	fmt.Println(c)
	c = a / b
	c = a + b
	c = a % b
	fmt.Println(c)
	d := 7.10 / 2.0
	fmt.Println(d)
	var i int32 = 5
	var j int64 = 7
	k := i + int32(j)
	fmt.Println(k)

	//Constants
	const aNew = 5 //-> implicit const
	//var new1 int64 -> exlplicit const wont work while assigning
	var new1 int16 = aNew
	fmt.Println(new1)
	const (
		aNew1 = "2"
		bNew
	)
	fmt.Println(bNew) // 2
	const (
		_      = iota
		tanza1 = 10 * iota //10
		tanza2             //20
		tanza3             //30
	)

	fmt.Print(tanza1, tanza2, tanza3)

	//Pointers
	sToPrint := "Hello gophers"
	p := &sToPrint
	fmt.Println(p)
	fmt.Println(*p)
	*p = "Help"
	fmt.Println(sToPrint) //Help

	//Reading input
	fmt.Println("Enter string:")
	in := bufio.NewReader(os.Stdin)
	sNew, _ := in.ReadString('\n')
	sNew = strings.TrimSpace(sNew)
	sNew = strings.ToUpper(sNew)
	fmt.Println(sNew)
}
