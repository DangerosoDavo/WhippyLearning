package main

import "fmt"

// ==========================================
// EXERCISE 1: Write a function
// Write a function called "multiply" that
// takes two ints and returns their product.
//
// func multiply(a, b int) int {
//     ...
// }
// ==========================================

// YOUR CODE HERE:

// ==========================================
// EXERCISE 3: Define a struct
// Create a struct called "Animal" with:
//   - Name (string)
//   - Sound (string)
//   - Legs (int)
// ==========================================

// YOUR CODE HERE:

func main() {
	// ==========================================
	// Test Exercise 1:
	// Uncomment these lines after writing multiply:
	//
	// fmt.Println(multiply(5, 3))   // should print 15
	// fmt.Println(multiply(10, 7))  // should print 70
	// ==========================================

	// ==========================================
	// EXERCISE 2: If/Else + Loop
	// Write a for loop that goes from 1 to 20.
	// For each number:
	//   - If divisible by 3, print "Fizz"
	//   - If divisible by 5, print "Buzz"
	//   - If divisible by both, print "FizzBuzz"
	//   - Otherwise, print the number
	//
	// Hint: divisible means remainder is 0
	//   if num % 3 == 0 { ... }
	// ==========================================

	fmt.Println("--- FizzBuzz ---")
	// YOUR CODE HERE:

	// ==========================================
	// EXERCISE 4: Use your struct
	// Create a slice of 3 Animal structs.
	// Loop through them and print:
	//   "The [Name] says [Sound] and has [Legs] legs"
	//
	// Example output:
	//   The Dog says Woof and has 4 legs
	//
	// Hint:
	//   animals := []Animal{
	//       {Name: "Dog", Sound: "Woof", Legs: 4},
	//       ...
	//   }
	// ==========================================

	fmt.Println("--- Animals ---")
	// YOUR CODE HERE:

	// ==========================================
	// EXERCISE 5 (BONUS): Multiple return values
	// Write a function called "stats" that takes
	// a slice of ints ([]int) and returns TWO
	// values: the smallest and the largest number.
	//
	// func stats(numbers []int) (int, int) { ... }
	//
	// Test it with:
	//   nums := []int{23, 7, 42, 3, 15, 99, 8}
	//   smallest, largest := stats(nums)
	//   fmt.Println("Smallest:", smallest, "Largest:", largest)
	// ==========================================

	fmt.Println("--- Stats ---")
	// YOUR CODE HERE:
}
