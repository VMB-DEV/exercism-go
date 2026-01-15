package resistorcolortrio

import (
	"fmt"
	"math"
)

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").

var colorCode = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

func Label(colors []string) string {
	if len(colors) < 3 {
		return ""
	}
	for _, color := range colors[:3] {
		if _, exist := colorCode[color]; !exist {
			return ""
		}
	}
	number := (colorCode[colors[0]]*10 + colorCode[colors[1]]) * int(math.Pow10(colorCode[colors[2]]))
	switch {
	case number >= 1_000_000_000:
		return fmt.Sprintf("%d gigaohms", number/1_000_000_000)
	case number >= 1_000_000:
		return fmt.Sprintf("%d megaohms", number/1_000_000)
	case number >= 1_000:
		return fmt.Sprintf("%d kiloohms", number/1_000)
	default:
		return fmt.Sprintf("%d ohms", number)
	}
}
