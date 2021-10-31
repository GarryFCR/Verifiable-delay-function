package src

import (
	"fmt"
	"math/big"
	"testing"
)

func TestVdf(test *testing.T) {

	N := Setup(5)
	x := Generate(N)
	t := 7
	list := []int{1, 2, 3}
	var i int
	for _, i = range list {
		y, proof, error := Solve(N, x, t, i)
		if !error {
			return
		}
		if !Verify(x, N, y, t, proof) {
			fmt.Println(i)
			test.Fatal("Verificaion failed")
		}

		proof[1] = *big.NewInt(111)

		if Verify(x, N, y, t, proof) {
			test.Fatal("Verificaion failed")
		}

	}

}
