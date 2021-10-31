package src

import (
	"fmt"
	"math"
	"math/big"
	"sort"
)

func Verify(x, N, y big.Int, t int, proof []big.Int) bool {

	T := int(math.Pow(2, float64(t)))

	quad := Quad_residues(N)
	signed_quad := Signed_quad_residues(N, quad)

	//check if the proof is in QRN+
	for _, j := range proof {

		pos := sort.Search(len(signed_quad), func(i int) bool {
			return j.Cmp(&signed_quad[i]) == 0 || j.Cmp(&signed_quad[i]) < 0
		})

		if pos == len(signed_quad) || j.Cmp(&signed_quad[pos]) != 0 {
			fmt.Println("proof not in QRN+")
			return false
		}
	}

	args := []big.Int{x, y}
	//check if x and y are in QRN+

	for _, k := range args {

		pos1 := sort.Search(len(signed_quad), func(i int) bool {
			return k.Cmp(&signed_quad[i]) == 0 || k.Cmp(&signed_quad[i]) < 0
		})

		if pos1 == len(signed_quad) || k.Cmp(&signed_quad[pos1]) != 0 {
			fmt.Println("x or y not in QRN+")
			return false
		}
	}

	var ui, yi, xi, ri big.Int
	var Ti int
	xi = *new(big.Int).Mul(&x, big.NewInt(1))
	yi = *new(big.Int).Mul(&y, big.NewInt(1))

	for i := 1; i <= t; i++ {

		ui = proof[i-1]
		Ti = T / int(math.Pow(2, float64(i-1)))
		ri = hashing(N, []big.Int{xi, *big.NewInt(int64(Ti)), yi, ui})
		//fmt.Printf("%v ", ri)

		xi = QRN_Exp_plus(xi, ri, N)
		xi.Mul(&xi, &ui)
		xi = QRN_Exp_plus(xi, *big.NewInt(1), N)

		temp := QRN_Exp_plus(ui, ri, N)
		yi = QRN_Exp_plus(*temp.Mul(&temp, &yi), *big.NewInt(1), N)

	}
	fmt.Println()
	xi = QRN_Exp_plus(xi, *big.NewInt(2), N)
	fmt.Println("here:", yi, xi)
	return yi.Cmp(&xi) == 0

}
