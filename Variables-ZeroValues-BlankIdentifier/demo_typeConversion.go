//Go dosn't allow type casting but it allows type conversion
//Type conversion is done using the T() syntax
//T() converts the value on the right hand side to the type T
//Type conversion is not allowed if the types are not compatible
//compatible types are types that have the same underlying type
//int and int32 are compatible types
//int and string are not compatible types

//you can convert float32 to float64 ,but you can't store a float64 in a float32 variable

package main

import "fmt"

func main() {

	fmt.Print("before conversion \n")

	var number float32 = 3.14
	fmt.Printf("number: %v , type: %T \n", number, number)
	fmt.Print("after conversion \n")
	var number2 float64 = float64(number)
	fmt.Printf("number2: %v , type: %T \n", number2, number2)

}
