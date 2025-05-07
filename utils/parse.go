package utils

import "strconv"

func ParseUint(s string) uint {
	id, _ := strconv.ParseUint(s, 10, 64)
	return uint(id)
}
