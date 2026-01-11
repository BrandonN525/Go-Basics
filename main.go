package main

import (
	"fmt"
	//"log"
	//"os"
)

func main() {
	fmt.Println("this is the beginning") // Print Statement
	fmt.Println()
	anotherFileTest() // Function from anotherfile.go
	fmt.Println()

	//Loops
	sum := 0
	//Go has no "while" loop, still use for
	for sum < 10 {
		sum += 1
		fmt.Println(sum)
	}

	fmt.Println() // Prints a blank line

	//Basic for loop
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(sum)
	}

	fmt.Println()

	//Using more modern Range

	for i := range 10 {
		sum += i
		fmt.Println(sum)
	}

	fmt.Println()

	//For range loop over slices

	r := []int{2, 4, 6, 8, 10}

	for i, v := range r {
		fmt.Println(i, v)
	}

	fmt.Println()

	/*Error Handling
	f, errf := os.Open("filename.txt")

	if err != nil {
		log.Fatal(err)
	}
	*/

	//Variables

	b := "this is a variable"
	fmt.Println(b)

	fmt.Println()

	var a string
	a = "this is a variable the long way"
	fmt.Println(a)

	fmt.Println()

	//Arrays and Slices

	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]

	//Use Printf instead of Println to insert variables into print statement
	fmt.Printf("Primes: %d\n", s)

	fmt.Println()

	//Maps

	m := make(map[string]int)

	m["ONE"] = 1
	m["TWO"] = 2
	m["THREE"] = 3

	fmt.Println(m)
}
