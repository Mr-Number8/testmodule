package testmodule

import (
	"fmt"
	"runtime"
)

// GoVersion v
func GoVersion() string {
	return fmt.Sprintf("Go version is <%s>", runtime.Version())
}

// Notag n
func Notag() string {
	return "get no tag"
}
