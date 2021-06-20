/*
; The Chudnovsky brothers' algorithm for calculating the digis of π.
; See http://en.wikipedia.org/wiki/Chudnovsky_algorithm
; 23 May 2013 - Scheme version
; 28 Feb 2015 - Go version
; Bakul Shah

(define (pi digits)
  (let* ((A 13591409)
	 (B 545140134)
	 (C 640320)
	 (C^3/24 (quotient (expt 640320 3) 24))
	 (D 12))
    (define (-1^n*g n g) (if (odd? n) (- g) g))
    (define (split m n)
      (if (= 1 (- n m))
	(let* ((6n (* 6 n)) (g (* (- 6n 5) (- (+ n n) 1) (- 6n 1))))
	  (list g (* C^3/24 (expt n 3)) (* (-1^n*g n g) (+ (* n B) A))))
	(let* ((mid (quotient (+ m n) 2))
	       (gpq1 (split m mid))
	       (gpq2 (split mid n))
	       (g1 (car gpq1)) (p1 (cadr gpq1)) (q1 (caddr gpq1))
	       (g2 (car gpq2)) (p2 (cadr gpq2)) (q2 (caddr gpq2)))
	  (list (* g1 g2) (* p1 p2) (+ (* q1 p2) (* q2 g1))))))
    (let* ((num-terms (inexact->exact (floor (+ 2 (/ digits 14.181647462)))))
	   (sqrt-C (integer-sqrt (* C (expt 100 digits))))
	   (gpq (split 0 num-terms))
	   (g (car gpq)) (p (cadr gpq)) (q (caddr gpq)))
      (quotient (* p C sqrt-C) (* D (+ q (* p A)))))))

; Compute time:
;	        (1)       (2)       (3)
; pi 100000     0.37      1.42      5.38
; pi 1000000    5.40    145.00     69.27
; pi 10000000  84.23  --------   2108.53
;
; (1) Gambit Scheme on 3.6GHz FX8150
; (2) Go 1.4        on 3.6Ghz FX8150
; (3) Gambit Scheme on 1.0Ghz ARMv7 (RaspberryPi 2)
; With Go I gave up on pi 1E7 after 7 hours.
*/
package main

import (
	"flag"
	"fmt"
	"math/big"
	"strconv"
	//"github.com/remyoudompheng/bigfft"
)

var (
	one = bigInt(1)
	two = bigInt(2)

	a          = bigInt(13591409)
	b          = bigInt(545140134)
	c          = bigInt(640320)
	cTo3Over24 = quo(expt(c, bigInt(3)), bigInt(24))
	d          = bigInt(12)
)

//  -1^n*g
func m1ToNmulG(n, g *big.Int) *big.Int {
	if odd(n) {
		return newInt().Neg(g)
	}
	return g
}

func split(m, n *big.Int) (g, p, q *big.Int) {
	if eqv(sub(n, m), one) {
		sixN := mul(bigInt(6), n)
		g := mul(mul(sub(sixN, bigInt(5)), sub(add(n, n), one)),
			sub(sixN, one))
		return g,
			mul(cTo3Over24, expt(n, bigInt(3))),
			mul(m1ToNmulG(n, g), add(mul(n, b), a))
	}
	mid := quo(add(m, n), two)
	g1, p1, q1 := split(m, mid)
	g2, p2, q2 := split(mid, n)
	return mul(g1, g2), mul(p1, p2), add(mul(q1, p2), mul(q2, g1))
}

func Π(digits int) *big.Int {
	terms := bigInt(2 + int(float64(digits)/14.181647462))
	sqrtC := isqrt(mul(c, expt(bigInt(100), bigInt(digits))))
	_, p, q := split(bigInt(0), terms)
	return quo(mul(p, mul(c, sqrtC)), mul(d, add(q, mul(p, a))))
}

func isqrt(n *big.Int) *big.Int {
	a, b := divMod(bigInt(n.BitLen()), two)
	x := expt(two, add(a, b))
	for {
		y := quo(add(x, quo(n, x)), two)
		if y.Cmp(x) >= 0 {
			return x
		}
		x = y
	}
}

func expt(x, y *big.Int) *big.Int { return newInt().Exp(x, y, nil) }
func quo(x, y *big.Int) *big.Int  { return newInt().Quo(x, y) }
func rem(x, y *big.Int) *big.Int  { return newInt().Rem(x, y) }
func sub(x, y *big.Int) *big.Int  { return newInt().Sub(x, y) }
func add(x, y *big.Int) *big.Int  { return newInt().Add(x, y) }

//func mul(x, y *big.Int)*big.Int { return bigfft.Mul(x, y) }
func mul(x, y *big.Int) *big.Int { return newInt().Mul(x, y) }
func eqv(x, y *big.Int) bool     { return x.Cmp(y) == 0 }
func divMod(n, m *big.Int) (a, b *big.Int) {
	return newInt().DivMod(n, m, newInt())
}

func bigInt(n int) *big.Int { return big.NewInt(int64(n)) }
func newInt() *big.Int      { return new(big.Int) }
func odd(n *big.Int) bool   { return eqv(rem(n, two), one) }

var flagv = flag.Bool("v", true, "verbose")

func main() {
	flag.Parse()
	dig := 100
	if flag.NArg() > 0 {
		dig, _ = strconv.Atoi(flag.Arg(0))
	}
	pi := Π(dig)
	if *flagv {
		fmt.Printf("%v\n", pi)
	}
}
