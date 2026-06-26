package utils

import (
	"math/rand"
	"time"
)

// Clamp returns value bounded by min and max
func Clamp(value, min, max int) int {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// Min returns the minimum of two values
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Max returns the maximum of two values
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// RandomInt returns a random integer between min and max (inclusive)
func RandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

// RandomBool returns a random boolean with a given percentage chance of true
func RandomBool(percentChance int) bool {
	return RandomInt(1, 100) <= percentChance
}

// CalculateCriticalChance calculates crit chance
func CalculateCriticalChance(baseCritChance int) bool {
	return RandomInt(1, 100) <= baseCritChance
}

// Abs returns the absolute value of an integer
func Abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

// PercentageOf calculates a percentage of a value
func PercentageOf(value, percentage int) int {
	return (value * percentage) / 100
}

// Average returns the average of two values
func Average(a, b int) int {
	return (a + b) / 2
}

// Scale scales a value based on a ratio
func Scale(value int, numerator, denominator int) int {
	if denominator == 0 {
		return 0
	}
	return (value * numerator) / denominator
}
