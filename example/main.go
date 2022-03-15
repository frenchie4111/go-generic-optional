package main

import (
	"fmt"

	"github.com/frenchie4111/go-generic-optional"
)

func main() {
	anOptional := opt.New[string]()

	if unwrapped, ok := anOptional.Get(); ok {
		fmt.Println("Unwrapped ", unwrapped)
	} else {
		fmt.Println("Could not unwrap")
	}

	hi := "hi"
	opt2 := opt.Of(hi)

	if unwrapped, ok := opt2.Get(); ok {
		fmt.Println("Unwrapped ", unwrapped)
	} else {
		fmt.Println("Could not unwrap")
	}

	thing := opt.If(opt2, func(ifunwrapped string) string {
		return "works"
	})
	fmt.Println("Handled ", thing.MustGet())
}
