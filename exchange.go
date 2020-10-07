package adtype

import (
	"strconv"
	"unicode/utf8"
)

func Str2Int(str string, defaultNumber int) int {
	if i, err := strconv.ParseInt(str, 0, 64); err != nil {
		return defaultNumber
	} else {
		return int(i)

	}
}
func Str2Long(str string, defaultNumber int64) int64 {
	if i, err := strconv.ParseInt(str, 0, 64); err != nil {
		return defaultNumber
	} else {
		return i
	}
}
func Str2Int32(str string, defaultNumber int32) int32 {
	if i, err := strconv.Atoi(str); err != nil {
		return defaultNumber
	} else {
		return int32(i)
	}
}
func Str2Double(str string, defaultNumber float64) float64 {
	if i, err := strconv.ParseFloat(str, 64); err != nil {
		return defaultNumber
	} else {
		return i
	}
}
func Str2Float(str string, defaultNumber float32) float32 {
	if i, err := strconv.ParseFloat(str, 32); err != nil {
		return defaultNumber
	} else {
		return float32(i)
	}
}

func Double2Str(x float64) string {
	return strconv.FormatFloat(x, 'f', -1, 64)
}

func Long2Str(x int64) string {
	return strconv.FormatInt(x, 10)
}
func Int2Str(x int) string {
	return strconv.Itoa(x)
}

func Any2Int32(in interface{}, defaultNumber int32) int32 {
	if n, ok := in.(int32); ok {
		return n
	}
	return defaultNumber

}
func Any2Int(in interface{}, defaultNumber int) int {
	if n, ok := in.(int); ok {
		return n
	}
	return defaultNumber

}
func Any2Long(in interface{}, defaultNumber int64) int64 {
	if n, ok := in.(int64); ok {
		return n
	}
	return defaultNumber
}

func Any2Float(in interface{}, defaultNumber float32) float32 {
	if n, ok := in.(float32); ok {
		return n
	}
	return defaultNumber
}
func Any2Double(in interface{}, defaultNumber float64) float64 {
	if n, ok := in.(float64); ok {
		return n
	}
	return defaultNumber
}

func Any2Str(any interface{}, defaultStr string) string {
	if str, ok := any.(string); ok {
		return str
	}
	return defaultStr
}

func Nano2Milli(nanoTimeStampStr string) string {
	length := utf8.RuneCountInString(nanoTimeStampStr)
	if length < 6 {
		return nanoTimeStampStr
	}
	return nanoTimeStampStr[:length-6]
}

func Milli2Nano(milliTimestamp string) string {
	return milliTimestamp + "000000"
}
