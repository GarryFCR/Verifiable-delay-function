package main

import (
	"fmt"

	"github.com/GarryFCR/VDF/src"
)

func main() {

	//run the vdf
	//generate the group
	var input int
	fmt.Println("Enter size of the RSA group in bits:")
	fmt.Scanln(&input)
	fmt.Println()
	N := src.Setup(input / 2)
	fmt.Println("N:", N.String())
	fmt.Println()

	//generate x randomly for y=x^2^t
	x := src.Generate(N)
	fmt.Println("x:", x.String())
	fmt.Println()

	//Start the sequential work
	//it generates the result i.e y
	//it generates the proof for the sequential work
	//
	//We assume T=2^t
	var t int
	fmt.Println("Enter the value of t (T=2^t):")
	fmt.Scanln(&t)
	//s is the parameter for precomputation(refer to the paper)
	// we choose s = {1,2,3} since computations get costly when s is larger
	var s int
	fmt.Println("Enter the value of s {1,2,3}:")
	fmt.Scanln(&s)
	if s > 3 || s < 0 {
		return
	}

	y, proof, error := src.Solve(N, x, t, s)
	if !error {
		return
	}
	fmt.Println()
	fmt.Println("y:", y.String())
	fmt.Println()
	fmt.Print("proof:")
	for i, _ := range proof {
		fmt.Printf("%v ", proof[i].String())
	}
	fmt.Println()
	//proof[0] = *big.NewInt(111)

	//Verify that y=x^2^t
	ver := src.Verify(x, N, y, t, proof)
	fmt.Println()
	fmt.Println("Verification:", ver)

}
