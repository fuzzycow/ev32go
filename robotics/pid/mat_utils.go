package pid

func absInt(i int) int {
	if i < 0 {
		return - i
	}
	return i
}

func absInt64(i int64) int64 {
	if i < 0 {
		return - i
	}
	return i
}


func normInt(want, abslimit int) int {
	switch {
	case want > 0 && want > abslimit:
		return abslimit
	case want < 0 && want < - abslimit:
		return -abslimit
	}
	return want
}

func normInt64(want, abslimit int64) int64 {
	switch {
	case want > 0 && want > abslimit:
		return abslimit
	case want < 0 && want < - abslimit:
		return -abslimit
	}
	return want
}

func maxAbsInt64(i1, i2 int64) int64 {
	if absInt64(i1) > absInt64(i2) {
		return i1
	} else {
		return i2
	}
}