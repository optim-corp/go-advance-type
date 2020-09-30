package ad_type

func MaxStr(strs ...string) string {
	if len(strs) == 0 {
		return ""
	}
	max := strs[0]
	for _, str := range strs {
		if str > max {
			max = str
		}
	}
	return max
}

func MinStr(strs ...string) string {
	if len(strs) == 0 {
		return ""
	}
	min := strs[0]
	for _, str := range strs {
		if str < min {
			min = str
		}
	}
	return min
}

func MaxInt64(args ...int64) int64 {
	if len(args) == 0 {
		return 0
	}
	max := args[0]
	for _, arg := range args {
		if arg > max {
			max = arg
		}
	}
	return max
}

func MinInt64(args ...int64) int64 {
	if len(args) == 0 {
		return 0
	}
	min := args[0]
	for _, arg := range args {
		if arg < min {
			min = arg
		}
	}
	return min
}
func MaxInt32(args ...int32) int32 {
	if len(args) == 0 {
		return 0
	}
	max := args[0]
	for _, arg := range args {
		if arg > max {
			max = arg
		}
	}
	return max
}

func MinInt32(args ...int32) int32 {
	if len(args) == 0 {
		return 0
	}
	min := args[0]
	for _, arg := range args {
		if arg < min {
			min = arg
		}
	}
	return min
}

func MaxInt(args ...int) int {
	if len(args) == 0 {
		return 0
	}
	max := args[0]
	for _, arg := range args {
		if arg > max {
			max = arg
		}
	}
	return max
}

func MinInt(args ...int) int {
	if len(args) == 0 {
		return 0
	}
	min := args[0]
	for _, arg := range args {
		if arg < min {
			min = arg
		}
	}
	return min
}
