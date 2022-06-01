package helper

import (
	"strconv"
)

func AddString(arg ...string) float64 {
	sum := float64(0)
	for _, s := range arg {
		f, _ := strconv.ParseFloat(s, 64)

		sum += f
	}

	return sum
}
