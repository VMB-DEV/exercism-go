package raindrops

import "strconv"

func Convert(number int) (ret string) {
	if number%3 == 0 {
		ret += "Pling"
	}
	if number%5 == 0 {
		ret += "Plang"
	}
	if number%7 == 0 {
		ret += "Plong"
	}
	if len(ret) == 0 {
		ret = strconv.Itoa(number)
	}
	return
}
