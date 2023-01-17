package main

import "github.com/khofesh/toolkit"

func main() {
	var tools toolkit.Tools

	tools.CreateDirIfNotExist("./test-dir")
}
