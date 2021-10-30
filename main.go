package main

import (
	"fmt"

	"github.com/GarryFCR/VDF/src"
)

func main() {
	N := src.Setup(5)
	fmt.Println("N:", N.String())
	fmt.Println()
	x := src.Generate(N)
	fmt.Println("x:", x)
	fmt.Println()

	y, proof := src.Solve(N, x, 3, 2)
	fmt.Println()

	fmt.Println("y:", y)
	fmt.Println()

	fmt.Println("proof:", proof)
	fmt.Println()

	ver := src.Verify(x, N, y, 3, proof)
	fmt.Println("Verification:", ver)

}
