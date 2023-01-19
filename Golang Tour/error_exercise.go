package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		var e ErrNegativeSqrt
		e = ErrNegativeSqrt(x);
		return 0, e
	}
	
	var z, prev float64
	z = 1
	prev = x + 1000000000
	
	for i := 0; i < 10; i++ {
		z -= (z*z - x)/(2*z)
		
		if math.Abs(z - prev) < 0.000000001 {
			return z, nil
		}
		prev = z
		fmt.Println(z)
	}
	
	return z, nil
}

// Ref: https://go.dev/tour/methods/20

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
