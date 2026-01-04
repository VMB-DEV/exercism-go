package leap

func IsLeapYear(year int) bool {
	switch {
	case year%4 == 0:
		return year%100 != 0 || year%400 == 0
	default:
		return false
	}
}
