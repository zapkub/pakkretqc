package perrors

import "fmt"

var (
	Unauthenticated = fmt.Errorf("Unauthorized")
)
