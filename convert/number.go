package convert

import (
	"golang.org/x/exp/constraints"
	"strings"
)

// ConvNum 数字转汉字：100 -> 一百
func ConvNum[T constraints.Integer](day T) string {
	if day < 0 {
		return ""
	}
	carry := [...]string{"", "", "", "", "", "万", "万", "万", "", "亿", "亿", "亿"}
	unit := [...]string{"", "十", "百", "千", "万", "十", "百", "千", "亿", "十", "百", "千"}
	num := [...]string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九", "十"}
	if day <= 10 {
		return num[day]
	}
	s := make([]T, 0, 3)
	for tempDay := day; tempDay > 0; {
		s = append(s, tempDay%10)
		tempDay = tempDay / 10
	}
	if len(s) > len(unit) {
		return ""
	}
	dataSlice := make([]string, 0, len(s))
	for k := len(s) - 1; k >= 0; k-- {
		if k == 0 && s[k] == 0 {
			if dataSlice[len(dataSlice)-1] == num[0] {
				dataSlice = dataSlice[0 : len(dataSlice)-1]
			}
			break
		}
		data := make([]string, 0, 4)
		if s[k] == 0 && dataSlice[len(dataSlice)-1] == num[0] {
			continue
		}
		if s[k] == 0 {
			data = append(data, num[s[k]])
		} else if carry[k] != "" {
			if len(dataSlice) > 1 && dataSlice[len(dataSlice)-1] == carry[k] {
				dataSlice = dataSlice[0 : len(dataSlice)-1]
			}
			data = append(data, num[s[k]], unit[k], carry[k])
		} else {
			if len(dataSlice) > 1 && dataSlice[len(dataSlice)-1] == unit[k] {
				dataSlice = dataSlice[0 : len(dataSlice)-1]
			}
			data = append(data, num[s[k]], unit[k])
		}
		dataSlice = append(dataSlice, data...)
	}
	return strings.Join(dataSlice, "")
}
