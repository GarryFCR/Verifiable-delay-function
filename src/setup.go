package src

import (
	"crypto/rand"
	"math/big"
)

func Setup(lamda int) big.Int {
	//lamda_rsa := 2*lamda
	var err bool = true
	var p, q *big.Int
	for err {
		var err1, err2 error
		p, err1 = safePrime(lamda)
		q, err2 = safePrime(lamda)
		err = !(err1 == nil && err2 == nil)
	}

	N := new(big.Int).Mul(p, q)
	return *N

}

func safePrime(bits int) (*big.Int, error) {

	//a := new(big.Int)
	one := big.NewInt(1)

	for {
		p, err := rand.Prime(rand.Reader, bits)

		if err != nil {

			return nil, err
		}
		// 2p+1
		a := new(big.Int).Lsh(p, 1)
		a = a.Add(a, one)
		if a.ProbablyPrime(20) {
			return a, nil
		}
	}
}
