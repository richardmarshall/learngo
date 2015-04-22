package main

import (
	"fmt"
	"math/rand"
)

func sorts() {
	fmt.Println("Sorty sort")
	funcs := []func([]int){
		Bubble,
		Insertion,
		Selection,
		Shell,
		Merge,
		Quick,
	}

	for _, fn := range funcs {
		input := rand.Perm(21)
		fmt.Println("Method: ", GetFunctionName(fn))
		fmt.Printf("Input:  %v\n", input)
		fn(input)
		fmt.Printf("Output: %v\n", input)
	}

}

func searches() {
	fmt.Println("Searchy search")
}

func graphs() {
	fmt.Println("Graphy graph")
}

func trees() {
	fmt.Println("Tree tree")
}

func main() {
	sorts()
	searches()
	graphs()
	trees()
}
