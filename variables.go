package main

import "fmt"

func main() {

	//Ways of declaring variables
	var a string = "apple" //method-1
	var b = 10             //method-2
	c := true              //method-3
	var d float32          //method-4
	fmt.Printf("the value of a :%v and type is %T \n", a, a)
	fmt.Printf("the value of b:%v and type is %T \n", b, b)
	fmt.Printf("the value of c :%v and type is %T\n", c, c)
	fmt.Printf("the value of d :%v and type is %T", d, d)

}
