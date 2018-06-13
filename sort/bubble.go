package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var list []int
	for i := 0; i < 100; i++ {
		s1 := rand.NewSource(time.Now().UnixNano())
		list = append(list, rand.New(s1).Intn(1000))
	}
	fmt.Printf("Ye old list: %v\n\n", list)
	bubble(list)
}

func bubble(list []int) {

	for i := 1; i < len(list)-1; i++ {

		if list[i-1] > list[i] {
			list[i], list[i-1] = list[i-1], list[i]
		}

	}
	fmt.Printf("The new list: %v\n\n", list)
}
