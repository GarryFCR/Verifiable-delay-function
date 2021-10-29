package src

import (
	"math"
	"math/big"
	"sort"
)

func Verify(x, N, y big.Int, t int, proof []big.Int) bool {

	T := int(math.Pow(2, float64(t)))

	quad := Quad_residues(N)
	signed_quad := Signed_quad_residues(N, quad)

	for _, j := range proof {
		if sort.Search(len(signed_quad), func(i int) bool {
			return j.Cmp(&signed_quad[i]) == 0
		}) == len(signed_quad) {
			return false
		}
	}

	args := []big.Int{x, y}
	for _, k := range args {
		if sort.Search(len(signed_quad), func(i int) bool {
			return k.Cmp(&signed_quad[i]) == 0
		}) == len(signed_quad) {
			return false
		}
	}

	var ui, yi, xi, ri big.Int
	var Ti int
	xi = *new(big.Int).Mul(&xi, big.NewInt(1))
	yi = *new(big.Int).Mul(&yi, big.NewInt(1))

	for i := 1; i <= t; i++ {

		ui = proof[i-1]
		Ti = T / int(math.Pow(2, float64(i-1)))
		ri = hashing(N, []big.Int{xi, *big.NewInt(int64(Ti)), yi, ui})

		xi = QRN_Exp_plus(xi, ri, N)
		xi.Mul(&xi, &ui)

		temp := QRN_Exp_plus(ui, ri, N)
		yi = QRN_Exp_plus(*temp.Mul(&temp, &yi), *big.NewInt(1), N)

	}

	xi = QRN_Exp_plus(xi, *big.NewInt(2), N)
	return yi.Cmp(&xi) == 0

}
