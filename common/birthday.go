package common

import "time"

// GetAge 根据生日获取年龄
func GetAge(birthday, format string) (int, error) {
	if birthday == "" {
		return 0, nil
	}
	d, err := time.Parse(format, birthday)
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
