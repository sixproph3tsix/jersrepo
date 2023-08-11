// pointer practice

// Pointers are like magical wands that can change a computers memory where the data is stored.  you can use pointers to change the data at that spot without copying everything


// declare the package main to kickoff the function and import fmt library from golang package
// call function main to initialize sequence
// set variable1 to a value that is a string and print
// set variable2 to the address of variable1 using Location Operator, &
//print variable2 and the dereferencing operator of variable2
//set variable2 value and print variable1

// package main

// import "fmt"

// func main() {
//     originalColor := "red"
//     fmt.Println("Original color:", originalColor)
    
//     colorPointer := &originalColor
//     fmt.Println("Color pointer:", colorPointer)
//     fmt.Println("Color at pointer:", *colorPointer)
    
//     *colorPointer = "blue"
//     fmt.Println("New color:", originalColor)
// }


// There can only be one toy (data) in the toy box (memory location) at a time.
// To ensure the toy doesn't get dirty (to avoid losing data), you ask your friend Temp to hold the toy while you work.
// You use the magic wand (dereferencing operator) to perform operations on the toy you're currently focused on (memory location pointed to by a pointer).
// Once you've completed the operation on the first toy (moved it to a new memory location), you can then move the second toy (variable/data) into the now-empty toy box (original memory location).
// This analogy beautifully captures the essence of using a temporary variable to safely swap values using pointers. It demonstrates the careful and deliberate process of data manipulation, ensuring that you don't lose any valuable information during the swapping process.




//declare your variables and their pointer addresses
// package main
// import ("fmt")

// func main() {

//     a := 5
//     b := 10

//     fmt.Println("The value of a is:", a)
//     fmt.Println("The value of b is:", b)
    
//     ptrA := &a
//     ptrB := &b 

//     fmt.Println("The memory address for a is:", ptrA)
//     fmt.Println("The memory address for b is", ptrB)

//     temp := *ptrA
//     *ptrA = *ptrB
//     *ptrB = temp

//     fmt.Println("The new value of a is:", a)
//     fmt.Println("The new value of b is:", b)


// }

package main 

import ("fmt")

func partOneSum(myArray, []int) {
    sum := 0
    for i := 0; i< len(myArray); i++ {

    }

}

func main() {
    myArray := [5] {3, 7, 12, 24, 18}
    partOneSum(myArray)
}

