package resistorcolortrio

import (
	"errors"
	"fmt"
	"strings"
)

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").

var colorData = map[string][]string{
	"black":  {"0", "", " ohms"},
	"brown":  {"1", "1", "0 ohms"},
	"red":    {"2", "2", " kiloohms"},
	"orange": {"3", "3", " kiloohms"},
	"yellow": {"4", "4", "0 kiloohms"},
	"green":  {"5", "5", " megaohms"},
	"blue":   {"6", "6", " megaohms"},
	"violet": {"7", "7", " gigaohms"},
	"grey":   {"8", "8", " gigaohms"},
	"white":  {"9", "9", " gigaohms"},
}

// ColorCode returns the resistance value of the given color.
func ColorData(color string) ([]string, error) {
	if data, exists := colorData[color]; exists {
		return data, nil
	} else {
		return data, errors.New(fmt.Sprintf("Color %s does not exist", color))
	}
}

func Label(colors []string) string {
	var elements []string
	for i, color := range colors[:3] {
		data, err := ColorData(color)
		if err != nil || i >= len(data) {
			return ""
		}
		elements = append(elements, data[i])
	}
	value := strings.Join(elements, "")
	if value[0] == '0' && value != "0 ohms" {
		value = value[1:]
	}
	return value
}
