package main

import (
	"fmt"
	"github.com/frenchie4111/go-generic-optional"
)

func main() {
	anOptional := opt.MakeOptional[string](nil)

	if unwrapped, err := anOptional.Unwrap(); err == nil {
		fmt.Println("Unwrapped ", unwrapped)
	} else {
		fmt.Println("Could not unwrap")
	}

	hi := "hi"
	opt2 := opt.MakeOptional(&hi)

	if unwrapped, err := opt2.Unwrap(); err == nil {
		fmt.Println("Unwrapped ", unwrapped)
	} else {
		fmt.Println("Could not unwrap")
	}

	thing := opt.If(opt2, func(ifunwrapped string) string {
		return "works"
	})
	fmt.Println("Handled ", thing.ForceUnwrap())
}
