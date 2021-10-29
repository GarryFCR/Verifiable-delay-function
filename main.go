package main

import (
	"fmt"

	"github.com/GarryFCR/VDF/src"
)

func main() {
	N := src.Setup(10)
	fmt.Println(N.String())

	set := src.Generate(N)
	fmt.Println(set)

	//src.Solve(N, *big.NewInt(125), 8, 2)

}
