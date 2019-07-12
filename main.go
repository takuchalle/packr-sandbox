package main

import (
	"fmt"
	"log"

	"github.com/gobuffalo/packr/v2"
)

func main() {
	box := packr.New("dataBox", "./data")

	s, err := box.FindString("setting.yaml")
	if err != nil {
	    log.Fatal(err)
	}
	fmt.Println(s)
}
