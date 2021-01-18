package fsutil

import (
	"os"
	"path"
)

var (
	wd, _  = os.Getwd()
	webdir = "web"
)

func PathFromWebDir(name string) string {
	return path.Join(wd, webdir, name)
}
