package kata

import (
	"fmt"
)

func SeriesSum(n int) string {
	if n == 0 {
		return "0.00"
	}
	if n == 1 {
		return "1.00"
	}

	var sum float64 = 1

	for i := 0; i < n-1; i++ {
		quotient := float64(4 + (i * 3))
		fmt.Println(quotient)
		sum += (1 / quotient)
	}

	return fmt.Sprintf("%.2f", sum)
}
