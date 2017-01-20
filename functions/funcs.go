package main

import (
	"errors"
	"fmt"
)

type Test func(string) error

func main() {
	testFunc := func(test string) error {
		if test == "" {
			return errors.New("test")
		}
		return nil
	}
	msg := testItOver(testFunc)
	fmt.Println(msg)
}

func testItOver(t Test) error {
	err := t("test")
	if err != nil {
		return err
	}
	return nil
}

func getItFrom() Test {
	testInsider := "this is a parent var"
	if testInsider != "" {
		return func(test string) error {
			if testInsider == "" {
				return errors.New("error, no string")
			}
			return nil
		}
	}
	return nil
}
