package adtype

func MaxStr(args ...string) string {
	if len(args) == 0 {
		return ""
	}
	max := args[0]
	for _, arg := range args {
		if arg > max {
			max = arg
		}
	}
	return max
}

func MinStr(args ...string) string {
	if len(args) == 0 {
		return ""
	}
	min := args[0]
	for _, arg := range args {
		if arg < min {
			min = arg
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

func MaxFloat32(args ...float32) float32 {
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

func MinFloat32(args ...float32) float32 {
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

func MaxFloat64(args ...float64) float64 {
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

func MinFloat64(args ...float64) float64 {
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
func SumStr(args ...string) string {
	result := ""
	for _, arg := range args {
		result += arg
	}
	return result
}
func SumInt(args ...int) int {
	result := 0
	for _, arg := range args {
		result += arg
	}
	return result
}
func SumInt32(args ...int32) int32 {
	result := int32(0)
	for _, arg := range args {
		result += arg
	}
	return result
}
func SumInt64(args ...int64) int64 {
	result := int64(0)
	for _, arg := range args {
		result += arg
	}
	return result
}

func SumFloat32(args ...float32) float32 {
	result := float32(0.0)
	for _, arg := range args {
		result += arg
	}
	return result
}
func SumFloat64(args ...float64) float64 {
	result := 0.0
	for _, arg := range args {
		result += arg
	}
	return result
}
func AverageInt(args ...int) int {
	v := 0
	for _, arg := range args {
		v += arg
	}
	if len(args) == 0 {
		return 0
	}
	return v / len(args)
}

func AverageInt32(args ...int32) int32 {
	v := int32(0)
	for _, arg := range args {
		v += arg
	}
	if len(args) == 0 {
		return 0
	}
	return v / int32(len(args))
}

func AverageInt64(args ...int64) int64 {
	v := int64(0)
	for _, arg := range args {
		v += arg
	}
	if len(args) == 0 {
		return 0
	}
	return v / int64(len(args))
}
func AverageFloat32(args ...float32) float32 {
	v := float32(0)
	for _, arg := range args {
		v += arg
	}
	if len(args) == 0 {
		return 0
	}
	return v / float32(len(args))
}
func AverageFloat64(args ...float64) float64 {
	v := float64(0)
	for _, arg := range args {
		v += arg
	}
	if len(args) == 0 {
		return 0
	}
	return v / float64(len(args))
}
