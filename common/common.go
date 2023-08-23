package common

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"time"
)

// Round 四舍五入,保留n位小数
func Round(f float64, n int) (float64, error) {
	x := fmt.Sprintf("%v", f)
	// math.Round(x*100) / 100
	xf, _, err := big.ParseFloat(x, 10, 256, big.ToNearestEven)
	if err != nil {
		return 0, err
	}
	p := math.Pow10(n)
	xfMulP, _ := new(big.Float).Mul(xf, big.NewFloat(p)).Float64()
	return math.Round(xfMulP) / p, nil
}

// RoundStr 四舍五入,保留n位小数
func RoundStr(x string, n int) (float64, error) {
	// math.Round(x*100) / 100
	xf, _, err := big.ParseFloat(x, 10, 256, big.ToNearestEven)
	if err != nil {
		return 0, err
	}
	p := math.Pow10(n)
	xfMulP, _ := new(big.Float).Mul(xf, big.NewFloat(p)).Float64()
	return math.Round(xfMulP) / p, nil
}

// GetAge 根据生日获取年龄
func GetAge(date, format string) (int, error) {
	if date == "" {
		return 0, nil
	}
	d, err := time.Parse(format, date)
	if err != nil {
		return 0, err
	}
	year := d.Year()
	month := int(d.Month())
	day := d.Day()
	age := 0
	if year <= 0 {
		age = 0
	}
	nowYear := time.Now().Year()
	nowMonth := int(time.Now().Month())
	nowDay := time.Now().Day()
	age = nowYear - year
	if nowMonth > month {
		age++
	} else if nowMonth < month {
		age--
	} else if nowMonth == month && nowDay >= day {
		age++
	} else if nowMonth == month && nowDay < day {
		age--
	}
	return age, nil
}

// IsCancel 判断上下文是否关闭
func IsCancel(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}
