package path

import (
	"path"
	"runtime"
)

func AbsPath() string {
	var absPath string
	_, filename, _, ok := runtime.Caller(1)
	if ok {
		absPath = path.Dir(filename)
	}
	return absPath
}
