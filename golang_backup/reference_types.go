package main

import (
	"fmt"
)

func main(){
	i := 42
	p := &i // reads i through pointer, assigning the memory address to p
	fmt.Println(*p) // will read the value of p by going to the memory addres
	i = 30
	fmt.Println(*p)
	*p = 40
	fmt.Println(i)
	*p = *p / 4
	fmt.Println(i)
}
