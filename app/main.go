package main

import (
	"fmt"

	"github.com/khofesh/toolkit"
)

func main() {
	var tools toolkit.Tools

	s := tools.RandomString(10)

	fmt.Println("random string:", s)
}
