# Go Optionals

This package is a tiny implementation of Optionals in go, using the new go1.18
generics.

## Example:

```go
package main

import (
	"fmt"
	"github.com/frenchie4111/go-generic-optional"
)

func main() {
	anOptional := opt.MakeOptional[string](nil)

	// Prints: "could not unwrap"
	if unwrapped, err := anOptional.Unwrap(); err == nil {
		fmt.Println("Unwrapped ", unwrapped)
	} else {
		fmt.Println("Could not unwrap")
	}

	hi := "hi"
	opt2 := opt.MakeOptional(&hi)

	// Prints: "Unwrapped  hi"
	if unwrapped, err := opt2.Unwrap(); err == nil {
		fmt.Println("Unwrapped ", unwrapped)
	} else {
		fmt.Println("Could not unwrap")
	}

	// Prints: "Unwrapped  works"
	thing := opt.If(opt2, func(ifunwrapped string) string {
		return "works"
	})
	fmt.Println("Handled ", thing.ForceUnwrap())
}
```
