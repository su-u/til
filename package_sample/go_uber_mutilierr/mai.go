package main

import (
	"errors"
	"fmt"

	"go.uber.org/multierr"
)

func main() {
	err := multierr.Combine(
		errors.New("call 1 failed"),
		nil,
		errors.New("call 2 failed"),
		// nil,
		errors.New("call 3 failed"),
	)

	fmt.Printf("%+v", err)
}
