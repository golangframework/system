package system

import (
	"runtime"
	"strings"
)

// Retrieves the name of the current package.
//
// # Returns
//
//	name string
//
// Name of the current package.
//
// # Example
//
//	package example
//	main {
//	    name := system.GetPackageName()
//	    fmt.Println("Package name:", name) // prints example
//	    }
//	}
func GetPackageName() (name string) {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		return "unknown"
	}
	name = runtime.FuncForPC(pc).Name()
	name = name[strings.LastIndex(name, "/")+1:]
	name = name[:strings.Index(name, ".")]
	return name
}
