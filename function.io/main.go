package main

import (
	"fmt"
	cal "function.io/operators" // Introduce cal prefix to use instead of package name ,since its long
)

func main()  {


	 fmt.Println("adding, sub , mul and dividing  two Integers 26 and 42 ")

	 add,sub,mul,div := cal.All(20,40)

 	 fmt.Println("out put is ", add , sub , mul, div)


}

