package main

import (
	"fmt"
)

func closureFunc() func(string) {
	additional := "from the caller"
	i := 0
	return func(message string) {
		i++
		fmt.Printf("%s, - %s - increment: %d \n", message, additional, i)
	}
}

func main() {
	printFunct := closureFunc()
	printFunct("test")
	printFunct("test")
	printFunct("test")
	printFunct("test")
}
