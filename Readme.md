# Go Optionals

This package is a tiny implementation of Optionals in go, using the new go1.18
generics.

Optionals are useful because sometimes errors are actually errors, and nil is a
valid response, but you want to retain type safety.

## Example:

```go
import (
	"fmt"
	"github.com/frenchie4111/go-generic-optional"
)

func getUser(userID: string) Optional[User], error {
	if !db.connected() {
		return opt.Empty[User](), fmt.Errorf("Failed to getUser Not connected to DB")
	}

	user, err := db.getUser()
	if err {
		return opt.Empty[User](), fmt.Errorf("Failed to getUser: %v", err)
	}

	return opt.Make(User)
}

func main() {
	user, err := getUser("user-id-1")

	if err != nil {
		// handle error cases
		return
	}

	opt.If(user, func(user User)) {
		fmt.Println("Found user", user.firstName)
	}
}
```
