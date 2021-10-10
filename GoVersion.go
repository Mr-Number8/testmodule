package testmodule

import (
	"fmt"
	"runtime"
)

func GoVersion() string {
	return fmt.Sprintf("Go version is <%s>", runtime.Version())
}
