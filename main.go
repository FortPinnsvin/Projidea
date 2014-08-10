package main

import (
	"fmt"
	"github.com/go-martini/martini"
)

func main() {
	m := martini.Classic()
	fmt.Println("Hello World!")
	m.Run()
}
