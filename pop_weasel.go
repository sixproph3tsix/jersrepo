package main

import (
	"fmt"
)

var UserInput int

func main() {
	fmt.Println("Please input an integer:")

	fmt.Scanln(&UserInput)
	//fmt.Println("Your number is",UserInput)

	if UserInput < 0 {
		popNeg(UserInput)
	} else if UserInput > 0 {
		popPlus(UserInput)
	} else {
		main()
	}
}

func popNeg(UserInput int) {
	for i := UserInput; i <= 0; i++ {
		if i%5 == 0 && i%3 == 0 {
			fmt.Println("min BANG!")
		} else if i%3 == 0 {
			fmt.Println("min pop")
		} else if i%5 == 0 {
			fmt.Println("min weasel")
		} else {
			fmt.Println(i)
		}
	}
}

func popPlus(UserInput int) {
	for i := 1; i <= UserInput; i++ {
		if i%5 == 0 && i%3 == 0 {
			fmt.Println("BANG!")
		} else if i%3 == 0 {
			fmt.Println("pop")
		} else if i%5 == 0 {
			fmt.Println("weasel")
		} else {
			fmt.Println(i)
		}
	}
}