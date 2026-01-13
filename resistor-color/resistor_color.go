package resistorcolor

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

// Colors returns the list of all colors.
func Colors() []string {
	var list []string
	for color, _ := range colorCode {
		list = append(list, color)
	}
	return list
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	if code, exists := colorCode[color]; exists {
		return code
	} else {
		return -1
	}
}
