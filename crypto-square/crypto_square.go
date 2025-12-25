package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

func Encode(pt string) string {
	if len(pt) == 0 {
		return ""
	}
	re := regexp.MustCompile("[^a-zA-Z0-9]+")
	pt = strings.ToLower(re.ReplaceAllString(pt, ""))
	ptLen := len(pt)
	c := int(math.Ceil(math.Sqrt(float64(ptLen))))
	r := int(math.Ceil(float64(ptLen) / float64(c)))
	if ptLen%c != 0 {
		pt += strings.Repeat(" ", c-ptLen%c)
	}
	targetLen := r * c
	if ptLen < targetLen {
		pt += strings.Repeat(" ", targetLen-len(pt))
	}
	var strBld strings.Builder
	for ci := range c {
		for ri := range r {
			strBld.WriteByte(pt[ri*c+ci])
		}
		if ci < c-1 {
			strBld.WriteByte(' ')
		}
	}
	return strBld.String()
}
