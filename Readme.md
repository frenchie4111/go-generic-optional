# Go Optionals

This package is a tiny implementation of Optionals in go, using the new go1.18
generics.

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
		panic(err)
	}

	opt.If(user, func(user User)) {
		fmt.Println("Found user", user.firstName)
	}
}
```
