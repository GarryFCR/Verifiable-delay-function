package src

import (
	"fmt"
	"math"
	"math/big"
)

func Solve(N, x big.Int, t, s int) {

	T := math.Pow(2, float64(t))
	precompute := math.Pow(2, float64(s))
	interval := big.NewInt(int64(T / precompute))
	var y big.Int = x
	var precompute_list []big.Int

	power := new(big.Int).Exp(big.NewInt(2), interval, nil)

	for i := 0; i < int(precompute); i++ {
		y = QRN_Exp_plus(y, *power, N)
		precompute_list = append(precompute_list, y)
	}
	fmt.Println(precompute_list)

}

func QRN_Exp_plus(x, y, N big.Int) big.Int {

	bound := new(big.Int).Div(new(big.Int).Sub(&N, big.NewInt(1)), big.NewInt(2))
	element := new(big.Int).Exp(&x, &y, &N)

	if element.Cmp(bound) == 1 {
		negative_element := new(big.Int).Neg(element)
		absolute_value := new(big.Int).Mod(negative_element, &N)
		return *absolute_value
	} else {
		return *element
	}

}

func ui(i, s int, N big.Int, precompute_list, r []big.Int) big.Int {

	var r_list, elements []big.Int
	var ui big.Int
	if i == 1 {
		elements = append(elements, precompute_list[3])
		r_list = append(r_list, *big.NewInt(1))
	} else if i == 2 {
		elements = append(elements, precompute_list[1], precompute_list[5])
		r_list = append(r_list, r[i-1], *big.NewInt(1))
	} else if i == 3 {
		elements = append(elements, precompute_list[0], precompute_list[4], precompute_list[2], precompute_list[6])
		r_list = append(r_list, *new(big.Int).Mul(&r[i-2], &r[i-3]), r[i-2], r[i-3], *big.NewInt(1))
	} else {
		return *big.NewInt(0)
	}

	for j := 0; j < int(math.Pow(2, float64(i-1))); j++ {
		power := QRN_Exp_plus(elements[j], r_list[j], N)
		X := *new(big.Int).Mul(&elements[j], &power)
		ui = QRN_Exp_plus(X, *big.NewInt(1), N)
	}

	return ui

}
