package main

import (
	"fmt"
	"math/big"

	"github.com/GarryFCR/VDF/src"
)

func main() {
	N := src.Setup(32)
	fmt.Println(N.String())

	set := src.Generate(15)
	fmt.Println(set)

	src.Solve(N, *big.NewInt(125), 8, 2)

}
