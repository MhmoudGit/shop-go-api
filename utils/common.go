package utils

import (
	"strconv"
)

func ToInt(s string) int {
	x, err := strconv.Atoi(s)
	if err != nil {
		panic(err.Error())
	}
	return x
}
