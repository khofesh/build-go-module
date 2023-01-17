package main

import (
	"log"

	"github.com/khofesh/toolkit"
)

func main() {
	toSlug := "hola mundo 666"

	var tools toolkit.Tools

	slugified, err := tools.Slugify(toSlug)
	if err != nil {
		log.Println(err)
	}

	log.Println(slugified)
}
