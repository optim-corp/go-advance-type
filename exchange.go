package ad_type

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

func Long2Str(x int64) string {
	return strconv.FormatInt(x, 10)
}
func Int2Str(x int) string {
	return strconv.Itoa(x)
}

func Any2Int32(in interface{}, defaultNumber int32) int32 {
	return Str2Int32(Any2Str(in, Int2Str(int(defaultNumber))), defaultNumber)

}

func Any2Int(in interface{}, defaultNumber int) int {
	return Str2Int(Any2Str(in, Int2Str(defaultNumber)), defaultNumber)

}
func Any2Long(in interface{}, defaultNumber int64) int64 {
	return Str2Long(Any2Str(in, Long2Str(defaultNumber)), defaultNumber)
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
