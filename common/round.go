package common

import (
	"fmt"
	"math"
	"math/big"
)

// Round 四舍五入,保留precision位小数(默认保留2位小数)
// f : The value to round
// precision : The optional number of decimal digits to round to.(default 2)
func Round(f float64, precision ...int) (float64, error) {
	width := 2
	if len(precision) > 0 {
		width = precision[0]
	}
	x := fmt.Sprintf("%v", f)
	// math.Round(x*100) / 100
	xf, _, err := big.ParseFloat(x, 10, 256, big.ToNearestEven)
	if err != nil {
		return 0, err
	}
	p := math.Pow10(width)
	xfMulP, _ := new(big.Float).Mul(xf, big.NewFloat(p)).Float64()
	return math.Round(xfMulP) / p, nil
}

// RoundStr 四舍五入,保留precision位小数(默认保留2位小数)
// f : The value to round
// precision : The optional number of decimal digits to round to.(default 2)
func RoundStr(x string, precision ...int) (float64, error) {
	width := 2
	if len(precision) > 0 {
		width = precision[0]
	}
	// math.Round(x*100) / 100
	xf, _, err := big.ParseFloat(x, 10, 256, big.ToNearestEven)
	if err != nil {
		return 0, err
	}
	p := math.Pow10(width)
	xfMulP, _ := new(big.Float).Mul(xf, big.NewFloat(p)).Float64()
	return math.Round(xfMulP) / p, nil
}
