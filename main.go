package main

import (
	"fmt"

	"github.com/GarryFCR/VDF/src"
)

func main() {
	N := src.Setup(32)
	fmt.Println(N.String())

	set := src.Generate(15)
	fmt.Println(set)

}
