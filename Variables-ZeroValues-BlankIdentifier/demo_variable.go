//In go variables are statically typed, which means their type can't change
//Variables are declared using the var keyword followed by the variable name and type
//Variables can be declared outside of a function, in which case they are package level variables
//Variables declared outside of a function are visible to all files in the package

//there is also shortdeclaration syntax for declaring variables
//this syntax can only be used inside a function
//the := operator is used for short declaration
//the type of the variable is automatically inferred from the value on the right hand side of the := operator
//the := operator can only be used for declaring new variables, it can't be used for assigning a new value to an existing variable
//in go we also have blank identifier, which is represented by the _ character
//the blank identifier can be used to discard values returned by a function
//the blank identifier can also be used to discard values while declaring variables

//go will throw error if you declare a variable and don't use it

//genral guidelines use var keyword if you need zero value for variable that is declared but not initialized
//use short declaration if you need to initialize the variable

package main

import "fmt"

func main() {

	fmt.Println("we are printing some values stored in variables through \n :shortdeclaration \n :declaration using var keyword\n")

	//short declaration

	name := "Amal krishna"
	fmt.Println("Short declaration:- name:", name)

	//declaration using var keyword

	var age int
	fmt.Println("printing with out initalization , age:", age)
	age = 22
	fmt.Println("printing after initalization , age:", age)

	//blank identifier

	name, place, _, age := "Amal krishna", "Kerala", "India", 22
	fmt.Println("name:", name, "\nplace:", place, "\nage:", age)

	/* Zero values for variabls in go

	   bool -> false
	   string -> ""
	   int -> 0
	   float -> 0.0
	   pointer -> nil

	*/

}
