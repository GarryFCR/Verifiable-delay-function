package src

import (
	//"fmt"
	"math/big"

	"github.com/cznic/sortutil"
)

/*func generate(N,T big.Int)(big.Int){


}*/

func Quad_residues(N int) []*big.Int {
	quad := make([]*big.Int, N)
	for i := 1; i < N; i++ {
		element := new(big.Int).Exp(big.NewInt(int64(i)), big.NewInt(2), big.NewInt(int64(N)))
		quad = append(quad[:], element)
	}

	unique_quad := make([]*big.Int, N)
	for i := 0; i < N; i++ {
		x := quad[i]
		n := sortutil.SearchBigInts(quad[:], x)
		if n >= 0 && n <= N {
			unique_quad = append(unique_quad, x)
		}
	}

	return unique_quad

}
