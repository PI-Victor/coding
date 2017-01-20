package main

import (
	"fmt"
)

func main() {
	newChan := make(chan []string)

	go func() {
		newString := []string{}
		for i := 0; i <= 100; i++ {
			newString = append(newString, string(i))
		}
		newChan <- newString
	}()
	fmt.Println(<-newChan)
	go func() {
		newStrings := []string{"1", "2", "3", "4", "5"}
		newChan <- newStrings
	}()

	fmt.Println(<-newChan)
}
