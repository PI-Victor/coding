package main

import (
	"context"
	"errors"
	"fmt"
	"math"
	"time"
)

const (
	errorCapMax = 10
)

func main() {
	ok, errors := retryFunction(failedFunction)

	if ok {
		if errors != nil {
			for _, err := range errors {
				fmt.Printf("Function executed with errors: %#v", err)
			}
			//NOTE: technically here we return true, but with errors not shown at this loglevel
		}
	}

	if !ok {
		// NOTE: we return failure followed by a generic error that masks the underlying errors.
		fmt.Println("Function execution failed with multiple errors, check the logs")
		fmt.Println(errors)
	}

	ok, errors = retryFunction(passFunction)

	if ok {
		if errors != nil {
			for _, err := range errors {
				fmt.Printf("Function executed with errors: %#v", err)
			}
			//NOTE: technically here we return true, but with errors not shown at this loglevel
		}
		fmt.Println("function executed successfully")
	}

	if !ok {
		// NOTE: we return failure followed by a generic error that masks the underlying errors.
		fmt.Println("Function execution failed with multiple errors, check the logs")
		fmt.Println(errors)
	}

}

func failedFunction(ctx context.Context) error {
	return errors.New("a new error")
}

func passFunction(ctx context.Context) error { return nil }

func retryFunction(f func(ctx context.Context) error) (bool, []error) {
	var errors []error

	for i := 1; i < errorCapMax+1; i++ {
		fmt.Printf("i is: %d\n", i)
		var (
			t    = time.Duration(math.Pow(2, float64(i)))
			wait = time.Millisecond * t
		)
		fmt.Println(wait)

		if err := f(context.Background()); err != nil {
			fmt.Printf("function execution failed retrying: %d and sleeping: %f seconds\n", i, wait.Seconds())
			errors = aggregateErrors(errors, err)
			time.Sleep(wait)
			fmt.Printf("Waited: %#v\n", wait)
			continue
		}
		return true, errors
	}
	return false, errors
}

func aggregateErrors(errors []error, err error) []error {
	if errors == nil {
		return append([]error{}, err)
	}

	return append(
		errors,
		err,
	)
}
