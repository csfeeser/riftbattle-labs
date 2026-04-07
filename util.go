package main

import "strconv"

func clamp(value, min, max int) int {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

func itoa(value int) string {
	return strconv.Itoa(value)
}
