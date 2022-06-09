package golang

import "io/fs"

func FilterGeneratedGoFiles(f fs.FileInfo) bool {
	if f.Name() == "generated.go" {
		return false
	}
	return true
}
