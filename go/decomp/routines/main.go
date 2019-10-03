package decomp

import (
	"fmt"
	"math/big"
	"sort"
	"strings"
)

type results struct {
	prime int
	count int
}

// Returns all primes between 2 and n
func primes(n int) []int {
	p := []int{}
	for i := 2; i <= n; i++ {
		if big.NewInt(int64(i)).ProbablyPrime(0) {
			p = append(p, i)
		}
	}
	return p
}

// Factor the prime fact into its primes
func primeFactorP(fact *big.Int, in <-chan int, out chan<- results) {
	q := <-in
	var dec results
	dec.prime = q
	z := big.NewInt(0)
	m := new(big.Int)

	for {
		fact, m = new(big.Int).DivMod(fact, big.NewInt(int64(q)), z)
		if m.Int64() != 0 {
			break
		}
		dec.count++
	}

	out <- dec
}

// decomp takes n and decomposes n! into it's prime factors
// returning a string of p^N values
// e. g. decomp(25) returns "2^22 * 3^10 * 5^6 * 7^3 * 11^2 * 13 * 17 * 19 * 23"
// for improved performance I use goroutines to run the prime factor calculations
func decomp(n int) string {
	dec := []string{}
	qs := primes(n)
	fact := new(big.Int).MulRange(1, int64(n))

	in := make(chan int)
	out := make(chan results, 100)

	for _, q := range qs {
		go primeFactorP(fact, in, out)
		in <- q
	}
	close(in)

	rs := []results{}
	for range qs {
		a := <-out
		rs = append(rs, a)
	}

	sort.Slice(rs, func(i, j int) bool {
		return rs[i].prime < rs[j].prime
	})

	for _, a := range rs {
		if a.count == 1 {
			dec = append(dec, fmt.Sprintf("%v", a.prime))
		} else {
			dec = append(dec, fmt.Sprintf("%v^%v", a.prime, a.count))
		}
	}
	return strings.TrimSpace(strings.Join(dec, " * "))
}
