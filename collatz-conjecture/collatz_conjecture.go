package collatzconjecture

import "errors"

func CollatzConjecture(n int) (step int, err error) {
	if n < 1 {
		return 0, errors.New("yolo")
	}
	for n > 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = n*3 + 1
		}
		step += 1
	}
	return
}
