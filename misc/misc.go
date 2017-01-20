package main

import (
	"fmt"
)

type testStruct struct {
	name string
}

type testInterface interface{}

func main() {
	test := &testStruct{
		name: "test",
	}

	copyStruct := *test
	copyStruct.name = "test11111"
	t := test.(*testInterface)
	fmt.Printf("%#v\n", t)
	fmt.Printf("%#v\n", &copyStruct)
}
