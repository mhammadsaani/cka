package main

import (
	"fmt"
	"math/cmplx"
)

func add(a int, b int) int {
	c := a + b
	return c
}

func add2(a, b int) (int, string)  {
	c := a + b
	return c, "google" 
	// multiple return values
}


// Another way to return values, by initializing variables instead of giving types. They will be returned by default
// return statement will be empty [Not a recommended way]

func split(sum int) (x, y int){
	x = sum * 20
	y = sum * 10
	return 
}


// func main()  {
// 	fmt.Println("Hello World", rand.Intn(10))
// 	fmt.Println(math.Pi)
// 	fmt.Println(add(2, 3))
// 	// deconstructing multiple values
// 	a, b := add2(2, 5)
// 	fmt.Println(a, b)
// 	fmt.Println(split(4))
// }


// Variables

// var variableName type outside a function inside a package [ the only way s]
// variables are initilized with default values, 
var a int
var b bool
// if variables of same type
var c, d, e float32 

// func main(){
// 	// if we know initial value we use variableName := 0 or var variableName = 0, but it is only inside function.
// 	// So, three ways
// 	//  - var variableName dataType [inside and outside function allowed, initilized to default value] use if you don't know the inital value
// 	//  - variableName :=  value [only inside function] use if you know the inital value

// 	i := 0
// 	fmt.Println(a, b, c, d, e, i)
// }


// --------------

var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 -1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)


// func  main()  {
// 	fmt.Printf("Type = %v value = %v\n", ToBe, ToBe)
// }

// %v => display value
// %T =? display data type
// %q => explicityly print the value in quote

// TYPE CONVERSION - No implicit conversion, only explicit convertion. Convert one data type to other dataype

// Example  - Not allowed, implicit conversion
// var a int32 = 10
// var b int64 = 20
// b = a

// func main(){
// 	var a int32 = 10
// 	var b int64 = 20
// 	b = a
// 	fmt.Println(b)
// }

// const Pi = 3.14 -> value cannot be changed later. Value is determined at compile domain. const can be either bool, int, string


// LOOPs -> Only For Loop, for can be converted to while and if nothing after for, it will be infinite loop

// func main(){
// 	// loop
// 	sum := 0
// 	i:=0
// 	for i < 5 {
// 		sum += i
// 		i+=1
// 	}
// 	fmt.Println(sum)
// }


// conditionals

// func main(){
// 	if v:=0; v<10{
// 		fmt.Println("Here")
// 	}else{
// 		fmt.Println("else")
// 	}
// }


// Defer -> code is pushed into  a stack [LIFO, FILO]

// Example-1
// func main(){
// 	defer fmt.Println("Later, into stack")
// 	defer fmt.Println("Deferred 2")

// 	fmt.Println("Normal Execution")
// }

// Example-2
func main(){
	fmt.Println("Couting")
	for i:=0; i<5; i++{
		defer fmt.Println(i)
	}
	fmt.Println("Done")
}

// Defer is used usually to do some tasks where we have something like
// We read a file and need to close later which we can forget so we use defer after reading, knowing the file will be closed later
// At end of function, let say we did the close, but a panic [execption] occurs, now the file is not closed which can cause memory 
// leaks. So, better way is to 
// open(file)
// defer close(file)
// database connection needs to be closed. 


