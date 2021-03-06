package src

import (
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

//Solve y=x^2^t
func Solve(N, x big.Int, t, s int) (big.Int, []big.Int, bool) {

	T := math.Pow(2, float64(t))
	precompute := math.Pow(2, float64(s))
	interval := T / precompute
	var y big.Int = x
	var precompute_list, proof []big.Int
	var error bool

	power := big.NewInt(int64(math.Pow(2, interval)))

	for i := 0; i < int(precompute); i++ {
		y = QRN_Exp_plus(y, *power, N)

		precompute_list = append(precompute_list, y)
	}

	proof, error = generate_proof(t, s, x, N, precompute_list)
	if !error {
		return *big.NewInt(0), []big.Int{}, error
	}
	return y, proof, error

}

//Exponent operation in signed quadratic residue group
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

//Generate the proof using precompute list
func generate_proof(t, s int, x, N big.Int, precompute_list []big.Int) ([]big.Int, bool) {

	var xi, ui, ri big.Int
	yi := precompute_list[len(precompute_list)-1]
	T := int(math.Pow(2, float64(t)))

	var Ti int

	var r_list []big.Int
	var proof []big.Int

	for i := 1; i <= s; i++ {

		var err1, err2, err3 bool
		ui, err1 = ui_new(i, s, N, precompute_list, r_list)

		xi, err2 = xi_new(i, s, x, N, precompute_list, r_list)

		yi, err3 = yi_new(i, s, N, precompute_list, r_list)

		if !(err1 && err2 && err3) {
			return []big.Int{}, false
		}
		proof = append(proof, ui)
		Ti = T / int(math.Pow(2, float64(i-1)))

		list := []big.Int{xi, *big.NewInt(int64(Ti)), yi, ui}
		ri = hashing(N, list)

		r_list = append(r_list, ri)
		//fmt.Println(ui, xi, yi)

	}

	//generating the remaining ui without using precompute list
	for j := s + 1; j <= t; j++ {

		xi = QRN_Exp_plus(xi, ri, N)
		xi = QRN_Exp_plus(*xi.Mul(&xi, &ui), *big.NewInt(1), N)

		temp := QRN_Exp_plus(ui, ri, N)
		yi = QRN_Exp_plus(*temp.Mul(&temp, &yi), *big.NewInt(1), N)

		power1 := math.Pow(2, float64(j))
		power2 := math.Pow(2, float64(T/int(power1)))
		power := *big.NewInt(int64(power2))

		ui = QRN_Exp_plus(xi, power, N)

		proof = append(proof, ui)
		Ti = T / int(math.Pow(2, float64(j-1)))

		list1 := []big.Int{xi, *big.NewInt(int64(Ti)), yi, ui}
		ri = hashing(N, list1)

		r_list = append(r_list, ri)
		//fmt.Println(ui, xi, yi)

	}
	//fmt.Println(r_list)
	return proof, true

}

//Used for hashing to generate ri
func hashing(N big.Int, list []big.Int) big.Int {

	X := new(big.Int)
	h := sha256.New()
	for _, y := range list {
		h.Write(y.Bytes())
	}
	hash := fmt.Sprintf("%x", h.Sum(nil))

	X.SetString(hash, 16)

	X.Mod(X, &N)

	return *X

}

//Generate ui
func ui_new(i, s int, N big.Int, precompute_list, r []big.Int) (big.Int, bool) {

	var r_list, elements []big.Int
	var ui big.Int

	if i == 1 {

		elements = append(elements, precompute_list[(len(precompute_list)/2)-1])
		r_list = append(r_list, *big.NewInt(1))

	} else if i == 2 {

		if s == 2 {
			elements = append(elements, precompute_list[0], precompute_list[s])

		} else {
			elements = append(elements, precompute_list[1], precompute_list[5])
		}

		r_list = append(r_list, r[0], *big.NewInt(1))

	} else if i == 3 {

		elements = append(elements, precompute_list[0], precompute_list[4], precompute_list[2], precompute_list[6])
		r_list = append(r_list, *new(big.Int).Mul(&r[0], &r[1]), r[1], r[0], *big.NewInt(1))

	} else {

		return *big.NewInt(0), false
	}

	temp := big.NewInt(1)
	for j := 0; j < int(math.Pow(2, float64(i-1))); j++ {

		power := QRN_Exp_plus(elements[j], r_list[j], N)
		temp.Mul(temp, &power)
		ui = QRN_Exp_plus(*temp, *big.NewInt(1), N)

	}

	return ui, true

}

//generate xi
func xi_new(i, s int, x, N big.Int, precompute_list, r []big.Int) (big.Int, bool) {

	var r_list, elements []big.Int
	var xi big.Int

	if i == 1 {

		elements = append(elements, x)
		r_list = append(r_list, *big.NewInt(1))

	} else if i == 2 {

		if s == 2 {
			elements = append(elements, x, precompute_list[1])
		} else {
			elements = append(elements, x, precompute_list[3])
		}

		r_list = append(r_list, r[0], *big.NewInt(1))

	} else if i == 3 {

		elements = append(elements, x, precompute_list[3], precompute_list[1], precompute_list[5])
		r_list = append(r_list, *new(big.Int).Mul(&r[0], &r[1]), r[1], r[0], *big.NewInt(1))

	} else {

		return *big.NewInt(0), false
	}

	temp := big.NewInt(1)
	for j := 0; j < int(math.Pow(2, float64(i-1))); j++ {

		power := QRN_Exp_plus(elements[j], r_list[j], N)
		temp.Mul(temp, &power)
		xi = QRN_Exp_plus(*temp, *big.NewInt(1), N)

	}

	return xi, true

}

//generate yi
func yi_new(i, s int, N big.Int, precompute_list, r []big.Int) (big.Int, bool) {

	var r_list, elements []big.Int
	var yi big.Int
	y := precompute_list[len(precompute_list)-1]

	if i == 1 {

		elements = append(elements, y)
		r_list = append(r_list, *big.NewInt(1))

	} else if i == 2 {

		if s == 2 {
			elements = append(elements, precompute_list[1], y)
		} else {
			elements = append(elements, precompute_list[3], y)
		}

		r_list = append(r_list, r[0], *big.NewInt(1))

	} else if i == 3 {

		elements = append(elements, precompute_list[1], precompute_list[5], precompute_list[3], precompute_list[7])
		r_list = append(r_list, *new(big.Int).Mul(&r[0], &r[1]), r[1], r[0], *big.NewInt(1))

	} else {

		return *big.NewInt(0), false
	}

	temp := big.NewInt(1)
	for j := 0; j < int(math.Pow(2, float64(i-1))); j++ {

		power := QRN_Exp_plus(elements[j], r_list[j], N)
		temp.Mul(temp, &power)
		yi = QRN_Exp_plus(*temp, *big.NewInt(1), N)

	}

	return yi, true

}
