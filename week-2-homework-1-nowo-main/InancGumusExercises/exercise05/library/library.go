package library

import "runtime"

func Version() string {
	return runtime.Version()
}
