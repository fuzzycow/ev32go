package utils

func SigFloat64(a float64) float64 {
	switch {
	case a < 0:
		return -1
	case a > 0:
		return +1
	}
	return 0
}
