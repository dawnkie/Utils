package singularity

func ConditionToInt(condition bool, Y, N int) int {
	if condition {
		return Y
	} else {
		return N
	}
}

func ConditionToBool(condition bool, Y, N bool) bool {
	if condition {
		return Y
	} else {
		return N
	}
}