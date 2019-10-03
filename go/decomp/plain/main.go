package decomp

import (
	"fmt"
	"math/big"
	"strings"
)

func primes(n int) []int {
	p := []int{}
	for i := 2; i <= n; i++ {
		if big.NewInt(int64(i)).ProbablyPrime(0) {
			p = append(p, i)
		}
	}
	return p
}

func decomp(n int) string {
	dec := []string{}
	z := big.NewInt(0)
	qs := primes(n)
	for _, qc := range qs {
		var c uint64
		fact := new(big.Int).MulRange(1, int64(n))
		m := new(big.Int)
		q := big.NewInt(int64(qc))

		for {
			fact, m = new(big.Int).DivMod(fact, q, z)
			if m.Int64() != 0 {
				break
			}
			c++
		}

		if c == 0 {
			break
		}
		if c == 1 {
			dec = append(dec, fmt.Sprintf("%v", q))
		} else {
			dec = append(dec, fmt.Sprintf("%v^%v", q, c))
		}
	}
	return strings.TrimSpace(strings.Join(dec, " * "))
}
