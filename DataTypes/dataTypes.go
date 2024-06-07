package main

import (
	"fmt"
	"maps"
)

func main() {
	fmt.Println("----Arrays------")
	var arr [4]string
	fmt.Println(arr)
	arr = [4]string{"a", "b", "c", "d"}
	fmt.Println(arr)
	arr2 := arr
	arr2[1] = "ajit"
	fmt.Println(arr2, arr) // [a ajit c d] [a b c d] Array Copies by value, data is not shared

	fmt.Println("----Slices-----")
	var s = []string{"a", "b", "c", "d"} //No fix size can shrink or expand
	s2 := s
	s2[1] = "ajit"
	fmt.Println(s2, s) // [a ajit c d] [a b c d] Slices are sharing the same data
	s3 := append(s, "home")
	fmt.Println(s3) //[a ajit c d home]

	fmt.Println("----Map-----")
	var m map[string]int
	fmt.Println(m)
	m = map[string]int{"foo": 1, "bar": 2}
	m["baz"] = 48
	fmt.Println(m)
	delete(m, "baz")
	fmt.Println(m)
	m["foo"] = 5
	fmt.Println(m)        //map[bar:2 foo:5]
	fmt.Println(m["new"]) //0
	newVar, ok := m["new"]
	fmt.Println(newVar, ok) //0 false

	fmt.Println("----Copying maps----")
	m2 := m
	m2["foo"] = 10
	fmt.Println(m)  //map[bar:2 foo:10]
	fmt.Println(m2) //map[bar:2 foo:10] //Map copies value by reference, Data is shared
	m3 := maps.Clone(m)
	m3["foo"] = 11
	fmt.Println(m)  //map[bar:2 foo:10]
	fmt.Println(m3) //map[bar:2 foo:11]
	// m==m3 //Maps are not comparable

	fmt.Println("----Map and Slice----")
	var ms map[string][]string
	fmt.Println(ms) //map[]
	ms = map[string][]string{"coffee": {"latte", "cappaciano"}, "tea": {"ginger", "assam"}}
	fmt.Println(ms) //map[coffe:[latte cappaciano] tea:[ginger assam]]

	ms["coffee"] = []string{"new"}
	fmt.Println(ms)
	v1, ok1 := ms["other"]
	fmt.Println(v1, ok1)

	fmt.Println("----Structs(Only heterogenous data type)-----")
	type student struct {
		name string
		roll int
	}
	var stud = student{name: "ajit", roll: 2}
	fmt.Println(stud) //{ajit 2}
	stud1 := stud
	stud.roll = 1
	fmt.Println(stud1, stud)   // {ajit 2} {ajit 1} structs are copied by value data is not shared
	fmt.Println(stud == stud1) // false structs are comparable
}
