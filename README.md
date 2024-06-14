Variables
Declaration:
    Variables in Go are declared using the var keyword.
    Example:
    var name string
    var age int
    var isStudent bool

    Short declaration (using :=) can be used within functions.
    name := "Alice"
    age := 21
    isStudent := true
Initialization:
    Variables can be declared and initialized in the same statement.

    var name string = "Alice"
    var age int = 21
    var isStudent bool = true

Type Inference:
    Go can infer the type of a variable based on the value assigned to it.
    var name = "Alice"
    var age = 21
    var isStudent = true

Simple Data Types
    Basic Types:
    Integers:
    Signed: int, int8, int16, int32, int64
    Unsigned: uint, uint8 (alias byte), uint16, uint32, uint64
    Floating Point:
    float32, float64
    Complex Numbers:
    complex64, complex128
    Boolean:
    bool (true or false)
    String:
    Immutable sequences of bytes.
    Example: var greeting string = "Hello, World!"

    Zero Values:
    Variables declared without an initial value are given a zero value.
    var i int          // 0
    var f float64      // 0.0
    var b bool         // false
    var s string       // ""

Aggregate Data Types
Arrays:
    Fixed-size sequences of elements of a single type.
    Array Copies by value, data is not shared
    Declaration:
    var arr [5]int
    arr[0] = 10
    fmt.Println(arr[0]) // Output: 10

Slices:
    Dynamically-sized, more common than arrays.
    Created using the built-in make function or slicing an array.
    Slices are sharing the same data

    var s []int = make([]int, 5)
    s = append(s, 10)

Maps:
    Key-value pairs.
    Created using the built-in make function.
    Map copies value by reference, Data is shared
    Maps are not comparable using assignment operator
    var m map[string]int = make(map[string]int)
    m["one"] = 1
    m["two"] = 2

Structs:
    Collections of fields.
    Used to group data together.
    structs are copied by value data is not shared

    type Person struct {
        Name string
        Age  int
    }

    var p Person
    p.Name = "Alice"
    p.Age = 21

Looping and Branching
Looping:

    For Loop:
    The only looping construct in Go.
    for i := 0; i < 10; i++ {
        fmt.Println(i)
    }

    For Range Loop:
    Iterates over elements in a variety of data structures (arrays, slices, maps, strings).
    arr := []int{1, 2, 3, 4, 5}
    for i, v := range arr {
        fmt.Println(i, v)
    }

    Infinite Loop:
    for {
        // Infinite loop
    }

Branching:

    If Statement:
    if age > 18 {
        fmt.Println("Adult")
    } else {
        fmt.Println("Minor")
    }

    If with Short Statement:
    if err := doSomething(); err != nil {
        fmt.Println("Error:", err)
    }

Switch Statement:
    A more readable way to execute one of several code blocks based on a value.
    switch day := 2; day {
    case 1:
        fmt.Println("Monday")
    case 2:
        fmt.Println("Tuesday")
    default:
        fmt.Println("Another day")
    }

    Switch without Expression:
    Useful for long if-else chains.
    switch {
    case age < 18:
        fmt.Println("Minor")
    case age >= 18:
        fmt.Println("Adult")
    }
