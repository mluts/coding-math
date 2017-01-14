package util

func Clampf(val, min, max float64) float64 {
	switch {
	case val < min:
		return min
	case val > max:
		return max
	default:
		return val
	}
}
