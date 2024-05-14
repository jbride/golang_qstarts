// The entry point of each go program is main.main, i.e. a function called main in a package called main. You have to provide such a main package.
package main

import "fmt"

/* reference
 *   https://gobyexample.com/closures
 */

// anonymous function returned as part of function "intSeq"
func intSeq() func() int {

	i := 0 // shorthand for declaring and initializing a variable.  This syntax is only available inside functions

	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq() // create a copy of anonymous function; scope of variables is specific to that anon function
	for x := 0; x < 5; x++ {
		fmt.Println((nextInt()))
	}

	newInts := intSeq() // create a new copy of anonymous function; scope of variables is specific to that anon function
	fmt.Println(newInts())
}
