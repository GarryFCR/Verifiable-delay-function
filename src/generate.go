package src

import (
	"fmt"
	"math/big"
	"math/rand"
	"sort"
	"time"
)

//generate x randomly from QRn+ i.e signed quadratic residue
func Generate(N int) big.Int {

	rand.Seed(time.Now().UnixNano())
	quad_residues := Quad_residues(N)
	signed_quad_residues := Signed_quad_residues(N, quad_residues)
	fmt.Println(signed_quad_residues)

	rand_index := rand.Intn(len(signed_quad_residues))
	x := signed_quad_residues[rand_index]

	return *x
}

//generate QRn group
func Quad_residues(N int) []*big.Int {

	quad := make([]*big.Int, N-1)
	for i := 1; i < N; i++ {
		element := new(big.Int).Exp(big.NewInt(int64(i)), big.NewInt(2), big.NewInt(int64(N)))
		quad[i-1] = element
	}
	unique_quad := dedupe(quad)
	fmt.Println(unique_quad)

	return unique_quad

}

//generate QRN+
func Signed_quad_residues(N int, quad_residues []*big.Int) []*big.Int {

	temp := (N - 1) / 2
	bound := big.NewInt(int64(temp))
	var signed_quad_residues []*big.Int
	for _, element := range quad_residues {
		if element.Cmp(bound) == 1 {
			negative_element := new(big.Int).Neg(element)
			absolute_value := new(big.Int).Mod(negative_element, big.NewInt(int64(N)))
			signed_quad_residues = append(signed_quad_residues[:], absolute_value)
		} else {
			signed_quad_residues = append(signed_quad_residues[:], element)
		}
	}
	result := dedupe(signed_quad_residues)
	return result
}

//remove duplicate from a set of big.Int after sorting
func dedupe(list []*big.Int) []*big.Int {

	sort.Slice(list, func(i, j int) bool {
		return list[i].Cmp(list[j]) == -1
	})

	unique_quad := []*big.Int{}
	unique_quad = append(unique_quad[:], list[0])

	for _, entry := range list {
		if entry.Cmp(unique_quad[len(unique_quad)-1]) != 0 {
			unique_quad = append(unique_quad[:], entry)
		}
	}
	return unique_quad
}
