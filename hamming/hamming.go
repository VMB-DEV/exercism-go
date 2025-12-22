package hamming

import "errors"

func Distance(a, b string) (d int, err error) {
	if len(a) != len(b) {
		return 0, errors.New("strings must be equal")
	}
	for i := range a {
		if a[i] != b[i] {
			d++
		}
	}
	return
}
