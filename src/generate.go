package src

import (
	"fmt"
	"math/big"
	"math/rand"
	"sort"
	"time"
)

func Generate(N int) big.Int {
	rand.Seed(time.Now().UnixNano())
	quad_residues := Quad_residues(N)
	rand_index := rand.Intn(len(quad_residues))
	x := quad_residues[rand_index]

	return *x
}

func Quad_residues(N int) []*big.Int {

	quad := make([]*big.Int, N-1)
	for i := 1; i < N; i++ {
		element := new(big.Int).Exp(big.NewInt(int64(i)), big.NewInt(2), big.NewInt(int64(N)))
		quad[i-1] = element
	}

	sort.Slice(quad, func(i, j int) bool {
		return quad[i].Cmp(quad[j]) == -1
	})
	fmt.Println(quad)

	unique_quad := []*big.Int{}
	unique_quad = append(unique_quad[:], quad[0])

	for _, entry := range quad {
		if entry.Cmp(unique_quad[len(unique_quad)-1]) != 0 {
			unique_quad = append(unique_quad[:], entry)
		}
	}

	fmt.Println(unique_quad)
	return unique_quad

}
