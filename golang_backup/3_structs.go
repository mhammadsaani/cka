package main

import (
	"fmt"
)



type nameOfStruct struct{ // definition
	X int
	y int
}

func main(){
	v := nameOfStruct{3, 4} // initilization 
	
	fmt.Println(v)

}


