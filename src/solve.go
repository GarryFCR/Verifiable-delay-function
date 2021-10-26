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
		y = QRN_plus(y, *power, N)
		precompute_list = append(precompute_list, y)
	}
	fmt.Println(precompute_list)

}

func QRN_plus(x, y, N big.Int) big.Int {

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
