package sumdecimal

import (
	"math/big"
	"strconv"
)

func SumDecimal(c int) int {
	if c < 1 {
		return 0
	}

	s := big.NewFloat(float64(c)).SetPrec(10000)
	square := big.NewFloat(2.0).SetPrec(10000)
	sqrt := square.Sqrt(s)

	asString := sqrt.Text('g', 10000)

	isMainPart := true

	var sum int
	var iterator int
	for _, v := range asString {
		if v == '.' {
			isMainPart = false
			continue
		}
		if isMainPart {
			continue
		}

		number, _ := strconv.Atoi(string(v))
		sum += number

		iterator++
		if iterator == 1000 {
			break
		}
	}
	return sum
}
